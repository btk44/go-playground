package transactions

import (
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

func NewTransactionHandler(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) GetTransactions(writer http.ResponseWriter, request *http.Request) {
	log.Println("server called")
	http.Error(writer, "test", http.StatusBadRequest)
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
