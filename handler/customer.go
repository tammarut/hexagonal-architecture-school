package handler

import (
	"code-bangkok/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	customerService service.CustomerService
}

func (this customerHandler) GetCustomers(writer http.ResponseWriter, req *http.Request) {
	customers, err := this.customerService.GetCustomers()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, err)
		return
	}
	fmt.Println("⏰Sleeping..")
	time.Sleep(500 * time.Millisecond)
	writer.Header().Set("content-type", "application/json")
	json.NewEncoder(writer).Encode(customers)
}

func (this customerHandler) GetACustomer(writer http.ResponseWriter, req *http.Request) {
	customerId := mux.Vars(req)["customerID"]
	customerID, err := strconv.Atoi(customerId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return
	}

	aCustomer, err := this.customerService.GetACustomer(customerID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, err)
		return
	}
	writer.Header().Set("content-type", "application/json")
	json.NewEncoder(writer).Encode(aCustomer)
	return
}

func NewCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{customerService: customerService}
}
