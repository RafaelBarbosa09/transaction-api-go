package handlers

import (
	"github.com/gin-gonic/gin"
	"transaction-api-go/application/services"
	"transaction-api-go/infrastructure/models"
)

type AccountHandler struct {
	AccountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accountService,
	}
}

func (handler *AccountHandler) HandleGetAllAccounts(ctx *gin.Context) {
	accounts, err := handler.AccountService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, accounts)
}

func (handler *AccountHandler) HandleCreateAccount(ctx *gin.Context) {
	var account models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdAccount, err := handler.AccountService.Create(account)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, createdAccount)
}
