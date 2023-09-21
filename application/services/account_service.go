package services

import (
	"transaction-api-go/domain/repositories"
	"transaction-api-go/infrastructure/models"
)

type AccountService struct {
	AccountRepository repositories.AccountRepository
}

func NewAccountService(accountRepository repositories.AccountRepository) *AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
	}
}

func (service *AccountService) GetAll() ([]models.Account, error) {
	return service.AccountRepository.GetAll()
}

func (service *AccountService) Create(account models.Account) (models.Account, error) {
	return service.AccountRepository.Create(account)
}
