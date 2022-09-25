package transactions

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Handler struct {
	logger     *log.Logger
	repository *Repository
}

const BaseUrl = "/transactions"

func NewTransactionHandler(logger *log.Logger) *Handler {
	return &Handler{logger: logger, repository: &Repository{}}
}

func (h *Handler) RegisterEndpoints(sm *http.ServeMux) {
	sm.HandleFunc(BaseUrl, h.handleRequest)
	sm.HandleFunc(BaseUrl+"/", h.handleRequest)
}

func (h *Handler) handleRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		if request.URL.Path == BaseUrl {
			h.getTransactions(writer, request)
			return
		}

		pathRegex := regexp.MustCompile("^" + BaseUrl + "/[0-9]+$")
		if pathRegex.Match([]byte(request.URL.Path)) {
			idRegex := regexp.MustCompile("[0-9]+$")

			matchSlice := idRegex.FindAllString(request.URL.Path, -1)
			if len(matchSlice) == 1 {
				id, err := strconv.Atoi(matchSlice[0])
				if err != nil {
					http.Error(writer, "this should never happen", http.StatusInternalServerError)
				}

				h.getTransactionById(writer, id)
				return
			}
		}

		http.Error(writer, "no path match.", http.StatusBadRequest)
		return
	}
}

func (h *Handler) getTransactions(writer http.ResponseWriter, request *http.Request) {
	h.logger.Println("Get transactions")
	transactions, err := h.repository.getTransactions()
	if err != nil {
		http.Error(writer, "database error", http.StatusInternalServerError)
		return
	}

	err = transactions.toJson(writer)
	if err != nil {
		http.Error(writer, "json encoding error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getTransactionById(writer http.ResponseWriter, id int) {
	h.logger.Println("Get one transaction")
	transaction, err := h.repository.getTransaction(id)
	if err != nil {
		http.Error(writer, "transaction not found", http.StatusNotFound)
		return
	}

	err = transaction.toJson(writer)
	if err != nil {
		http.Error(writer, "json encoding error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddTransaction(writer http.ResponseWriter, request *http.Request) {
	h.logger.Println("Add transaction")

	transaction := &Transaction{}
	err := transaction.fromJson(request.Body)
	if err != nil {
		http.Error(writer, "json decoding error", http.StatusBadRequest)
		return
	}

	h.repository.addTransaction(transaction)
}
