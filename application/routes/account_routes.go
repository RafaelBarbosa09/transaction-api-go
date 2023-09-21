package routes

import (
	"github.com/gin-gonic/gin"
	"transaction-api-go/application/handlers"
)

func LoadAccountRoutes(router *gin.Engine, accountHandler *handlers.AccountHandler) {
	accountGroup := router.Group("/accounts")
	{
		accountGroup.GET("", accountHandler.HandleGetAllAccounts)
		accountGroup.POST("", accountHandler.HandleCreateAccount)
	}
}
