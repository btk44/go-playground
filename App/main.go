package main

import (
	//transactions "App/transactions"

	"App/transactions"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "api", log.LstdFlags)

	sm := http.NewServeMux()

	transactionHandler := transactions.NewTransactionHandler(logger)
	transactionHandler.RegisterEndpoints(sm)

	s := &http.Server{
		Addr:         ":6002",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Fatal(s.ListenAndServe())
	}()

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, os.Interrupt)

	shutDownSignal := <-osChan
	logger.Println("\n Shutting down the server. Reason: ", shutDownSignal)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
