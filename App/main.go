package main

import (
	//transactions "App/transactions"

	"App/transactions"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "api", log.LstdFlags)
	transactionHandler := transactions.NewTransactionHandler(logger)

	http.HandleFunc("/", transactionHandler.GetTransactions)

	http.ListenAndServe(":6001", nil)
}
