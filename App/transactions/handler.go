package transactions

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

func NewTransactionHandler(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) RegisterEndpoints(sm *http.ServeMux) {
	sm.HandleFunc("/transactions", h.GetTransactions)
}

func (h *Handler) GetTransactions(writer http.ResponseWriter, request *http.Request) {
	result, err := json.Marshal(transactionsData)
	if err != nil {
		http.Error(writer, "oops! json issue", http.StatusInternalServerError)
		return
	}

	writer.Write(result)
}

// func GetTransactions(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, transactionsData)
// }

// func GetTransactionById(context *gin.Context) {
// 	id, err := strconv.Atoi(context.Param("id"))

// 	if err != nil {
// 		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id provided."})
// 		return
// 	}

// 	for i, trans := range transactionsData {
// 		if trans.Id == id {
// 			context.IndentedJSON(http.StatusOK, transactionsData[i])
// 			return
// 		}
// 	}

// 	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transactionnot found."})
// }
