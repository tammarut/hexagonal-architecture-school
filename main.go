package main

import (
	"code-bangkok/handler"
	"code-bangkok/repository"
	"code-bangkok/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initializeDatabase() *sqlx.DB {
	datasourceName := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.databaseName"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), datasourceName)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	initTimeZone()
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()

	db := initializeDatabase()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	// _ = customerRepository
	// customerRepositoryStub := repository.NewCustomerRepositoryStub()
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetACustomer).Methods(http.MethodGet)

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	log.Printf("Start server at port %d\n", viper.GetInt("app.port"))
	http.ListenAndServe(port, router)
}
