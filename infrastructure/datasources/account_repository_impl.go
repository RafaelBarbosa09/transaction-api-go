package datasources

import (
	"github.com/jmoiron/sqlx"
	"transaction-api-go/infrastructure/models"
)

type AccountRepositoryImpl struct {
	db *sqlx.DB
}

func NewAccountRepositoryImpl(db *sqlx.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		db,
	}
}

func (repository *AccountRepositoryImpl) Create(account models.Account) (models.Account, error) {
	var id int64
	if err := repository.db.QueryRow(
		"INSERT INTO account (account_holder, active_card, available_limit) VALUES ($1, $2, $3) RETURNING id",
		account.AccountHolder,
		account.ActiveCard,
		account.AvailableLimit,
	).Scan(&id); err != nil {
		return models.Account{}, err
	}

	account.ID = id
	return account, nil
}

func (repository *AccountRepositoryImpl) GetAll() ([]models.Account, error) {
	var accounts []models.Account
	err := repository.db.Select(&accounts, "SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
