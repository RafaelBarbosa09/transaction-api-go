package di

import (
	"github.com/jmoiron/sqlx"
	"transaction-api-go/application/handlers"
	"transaction-api-go/application/services"
	"transaction-api-go/infrastructure/datasources"
)

type Container struct {
	AccountHandler *handlers.AccountHandler
}

func NewContainer(db *sqlx.DB) *Container {

	accountRepository := datasources.NewAccountRepositoryImpl(db)

	accountService := services.NewAccountService(accountRepository)

	accountHandler := handlers.NewAccountHandler(*accountService)

	return &Container{
		AccountHandler: accountHandler,
	}
}
