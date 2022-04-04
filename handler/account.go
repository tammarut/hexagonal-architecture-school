package handler

import (
	"code-bangkok/errs"
	"code-bangkok/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountSrv service.AccountService) accountHandler {
	return accountHandler{accountService: accountSrv}
}

func (this accountHandler) NewAccount(writer http.ResponseWriter, req *http.Request) {
	customerId := mux.Vars(req)["customerID"]
	customerID, err := strconv.Atoi(customerId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return
	}

	reqContentType := req.Header.Get("Content-Type")
	if reqContentType != "application/json" {
		handleError(writer, errs.NewValidationError("Request body is wrong format"))
		return
	}

	request := service.AccountRequest{}
	decodeErr := json.NewDecoder(req.Body).Decode(&request)
	if decodeErr != nil {
		handleError(writer, errs.NewValidationError("Request body is wrong format"))
	}

	response, err := this.accountService.NewAccount(customerID, request)
	if err != nil {
		handleError(writer, err)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func (this accountHandler) GetAccounts(writer http.ResponseWriter, req *http.Request) {
	customerId := mux.Vars(req)["customerID"]
	customerID, err := strconv.Atoi(customerId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return
	}

	responses, err := this.accountService.GetAccounts(customerID)
	if err != nil {
		handleError(writer, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(responses)
	return
}
