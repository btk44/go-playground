package transactions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTransactions(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, transactionsData)
}

func GetTransactionById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id provided."})
		return
	}

	for i, trans := range transactionsData {
		if trans.Id == id {
			context.IndentedJSON(http.StatusOK, transactionsData[i])
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transactionnot found."})
}
