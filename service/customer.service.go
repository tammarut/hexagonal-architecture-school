package service

import (
	"code-bangkok/repository"
	"database/sql"
	"errors"
	"log"
)

func buildResponses(customers []repository.Customer) []CustomerResponse {
	var customerResponses []CustomerResponse
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses
}

type customerService struct {
	customerRepo repository.CustomerRepository
}

func (this customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := this.customerRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	customerResponses := buildResponses(customers)
	return customerResponses, nil
}

func (this customerService) GetACustomer(customerId int) (*CustomerResponse, error) {
	aCustomer, err := this.customerRepo.GetById(customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Customer not found❌")
		}
		log.Println(err)
		return nil, err
	}
	customerResponse := CustomerResponse{
		CustomerID: aCustomer.CustomerID,
		Name:       aCustomer.Name,
		Status:     aCustomer.Status,
	}
	return &customerResponse, nil
}

// NewCustomerService is like a constructor to initialize an object
func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return customerService{customerRepo: customerRepo}
}
