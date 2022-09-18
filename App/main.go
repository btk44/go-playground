package main

import (
	transactions "App/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/transactions", transactions.GetTransactions)
	router.GET("/transactions/:id", transactions.GetTransactionById)
	router.Run("localhost:8080")
}
