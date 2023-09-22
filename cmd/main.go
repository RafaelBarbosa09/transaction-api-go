package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"transaction-api-go/application/routes"
	"transaction-api-go/config"
	"transaction-api-go/di"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := config.InitializeDataBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	container := di.NewContainer(db)

	router := gin.Default()

	routes.LoadAccountRoutes(router, container.AccountHandler)

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
