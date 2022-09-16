package transactions

import (
	common "App/common"
	"fmt"
	"time"
)

type Transaction struct {
	Id        int
	Date      time.Time
	AccountId int
	Amount    float32
	Payee     string
	Category  int
	Comment   string
}

func Add(transaction Transaction) common.Result[bool] {
	return common.GetFailedResult[bool]("error message")
}

func Delete(transactionId int) common.Result[bool] {
	return common.GetSuccessResult(true)
}

func Update(transaction *Transaction) error {
	transaction.AccountId = 1

	return fmt.Errorf("there was an error")
}

func Update2(transaction Transaction) error {
	transaction.AccountId = 7
	return nil
}
