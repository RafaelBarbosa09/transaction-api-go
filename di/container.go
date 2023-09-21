package di

import (
	"transaction-api-go/application/handlers"
	"transaction-api-go/application/services"
	"transaction-api-go/infrastructure/datasources"
)

type Container struct {
	AccountHandler *handlers.AccountHandler
}

func NewContainer() *Container {
	accountRepository := datasources.NewAccountRepository()

	accountService := services.NewAccountService(accountRepository)

	accountHandler := handlers.NewAccountHandler(*accountService)

	return &Container{
		AccountHandler: accountHandler,
	}
}
