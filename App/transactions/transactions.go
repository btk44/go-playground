package transactions

import (
	common "App/common"
	"time"
)

type Transaction struct {
	id        int
	date      time.Time
	accountId int
	amount    float32
	payee     string
	category  int
	comment   string
}

func Add(transaction Transaction) common.Result[bool] {
	return common.GetFailedResult[bool]("dupa")
}

func Delete(transactionId int) common.Result[bool] {
	return common.GetSuccessResult(true)
}
