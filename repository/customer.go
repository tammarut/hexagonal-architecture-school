package repository

import "errors"

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(customerId int) (*Customer, error)
}

type customerRepositoryStub struct {
	customers []Customer
}

func (this customerRepositoryStub) GetAll() ([]Customer, error) {
	return this.customers, nil
}

func (this customerRepositoryStub) GetById(customerId int) (*Customer, error) {
	for _, customer := range this.customers {
		if customer.CustomerID == customerId {
			return &customer, nil
		}
	}

	notFoundErr := errors.New("Not found this customer id")
	return nil, notFoundErr
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashi", City: "New York", DateOfBirth: "01/09/2020", ZipCode: "71999", Status: 1},
		{CustomerID: 1002, Name: "Jo", City: "LA", DateOfBirth: "03/04/2021", ZipCode: "71999", Status: 1},
	}

	return customerRepositoryStub{customers: customers}
}
