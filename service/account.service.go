package service

import (
	"code-bangkok/errs"
	"code-bangkok/logger"
	"code-bangkok/repository"
	"time"
)

const (
	INACTIVE int = 0
	ACTIVE       = 1
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func (this accountService) NewAccount(customerID int, metaData AccountRequest) (*AccountResponse, error) {
	// Validate metaData

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-2 15:04:05"),
		AccountType: metaData.AccountType,
		Amount:      metaData.Amount,
		Status:      ACTIVE,
	}

	newAcc, err := this.accountRepo.CreateNew(account)
	if err != nil {
		logger.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (this accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := this.accountRepo.GetAll(customerID)
	if err != nil {
		logger.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := this.toDTO(accounts)

	return responses, nil
}

func (this accountService) toDTO(accounts []repository.Account) []AccountResponse {
	responses := []AccountResponse{}
	for _, account := range accounts {
		currentAccount := AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		}
		responses = append(responses, currentAccount)
	}
	return responses
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}
