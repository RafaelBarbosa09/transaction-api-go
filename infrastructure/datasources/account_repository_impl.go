package datasources

import "transaction-api-go/infrastructure/models"

type AccountRepositoryImpl struct {
	accounts []models.Account
}

func NewAccountRepository() *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		accounts: []models.Account{},
	}
}

func (repository *AccountRepositoryImpl) GetAll() ([]models.Account, error) {
	return repository.accounts, nil
}

func (repository *AccountRepositoryImpl) Create(account models.Account) (models.Account, error) {
	repository.accounts = append(repository.accounts, account)
	return account, nil
}
