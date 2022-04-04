package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryAdapter struct {
	db *sqlx.DB
}

func (this accountRepositoryAdapter) CreateNew(acc Account) (*Account, error) {
	command := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	record, err := this.db.Exec(command, acc.CustomerID, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		return nil, err
	}

	id, err := record.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)

	return &acc, nil
}

func (this accountRepositoryAdapter) GetAll(customerID int) ([]Account, error) {
	query := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE customer=?"
	var accounts []Account
	queryErr := this.db.Select(&accounts, query, customerID)
	if queryErr != nil {
		return nil, queryErr
	}

	return accounts, nil
}

func NewAccountRepositoryDB(sqlxDB *sqlx.DB) AccountRepository {
	return accountRepositoryAdapter{db: sqlxDB}
}
