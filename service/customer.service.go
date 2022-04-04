package service

import (
	"code-bangkok/errs"
	"code-bangkok/logger"
	"code-bangkok/repository"
	"database/sql"
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
		logger.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customerResponses := buildResponses(customers)
	return customerResponses, nil
}

func (this customerService) GetACustomer(customerId int) (*CustomerResponse, error) {
	aCustomer, err := this.customerRepo.GetById(customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found❌")
		}
		logger.Error(err)
		return nil, errs.NewUnexpectedError()
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
