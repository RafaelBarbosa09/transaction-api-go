package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transaction-api-go/application/routes"
	"transaction-api-go/di"
)

func main() {
	container := di.NewContainer()

	router := gin.Default()

	routes.LoadAccountRoutes(router, container.AccountHandler)

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
