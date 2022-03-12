package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepositoryAdapter struct {
	db *sqlx.DB
}

func (this customerRepositoryAdapter) GetAll() ([]Customer, error) {
	var customers []Customer
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
	err := this.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (this customerRepositoryAdapter) GetById(customerId int) (*Customer, error) {
	var customer Customer
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id=?"
	err := this.db.Get(&customer, query, customerId)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryAdapter{db: db}
}
