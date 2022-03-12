package main

import (
	"code-bangkok/handler"
	"code-bangkok/repository"
	"code-bangkok/service"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()
	db, err := sqlx.Open("mysql", "Arima_kishou0:My_secrete_passw0rd@tcp(localhost:3306)/banking?parseTime=true")
	if err != nil {
		panic(err)
	}
	customerRepository := repository.NewCustomerRepositoryDB(db)
	_ = customerRepository
	customerRepositoryStub := repository.NewCustomerRepositoryStub()
	customerService := service.NewCustomerService(customerRepositoryStub)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetACustomer).Methods(http.MethodGet)

	http.ListenAndServe(":9000", router)
}
