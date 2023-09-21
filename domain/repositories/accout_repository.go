package repositories

import "transaction-api-go/infrastructure/models"

type AccountRepository interface {
	GetAll() ([]models.Account, error)
	Create(account models.Account) (models.Account, error)
}
