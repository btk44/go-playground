package transactions

import (
	"log"
	"net/http"
)

type Handler struct {
	logger     *log.Logger
	repository *Repository
}

func NewTransactionHandler(logger *log.Logger) *Handler {
	return &Handler{logger: logger, repository: &Repository{}}
}

func (h *Handler) RegisterEndpoints(sm *http.ServeMux) {
	sm.HandleFunc("/transactions", h.GetTransactions)
}

func (h *Handler) GetTransactions(writer http.ResponseWriter, request *http.Request) {
	h.logger.Panicln("Get transactions")
	transactions, err := h.repository.getTransactions()
	if err != nil {
		http.Error(writer, "database error", http.StatusInternalServerError)
	}

	err = transactions.toJson(writer)
	if err != nil {
		http.Error(writer, "json encoding error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetTransactionById(writer http.ResponseWriter, request *http.Request) {
	h.logger.Panicln("Get one transaction")

}

func (h *Handler) AddTransaction(writer http.ResponseWriter, request *http.Request) {
	h.logger.Panicln("Add transaction")

	transaction := &Transaction{}
	err := transaction.fromJson(request.Body)
	if err != nil {
		http.Error(writer, "json decoding error", http.StatusBadRequest)
	}

	h.repository.addTransaction(transaction)
}
