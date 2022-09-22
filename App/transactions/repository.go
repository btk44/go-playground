package transactions

import (
	"time"
)

type Repository struct {
}

func (r *Repository) getTransactions() (Transactions, error) {
	return transactionsData, nil
}

func (r *Repository) getTransaction(id int) *Transaction {
	// transactionIndex := Transactions.IndexFunc(transactionsData, func(t *Transaction) bool {
	// 	return t.Id == id
	// })

	// if transactionIndex == -1 {
	// 	return nil
	// }

	// return transactionsData[transactionIndex]
	return nil
}

func (r *Repository) addTransaction(transaction *Transaction) {
	lastIndex := len(transactionsData) - 1
	transaction.Id = transactionsData[lastIndex].Id + 1
	transactionsData = append(transactionsData, transaction)
}

func (r *Repository) updateTransaction(transaction *Transaction) {
	//oldTransaction := r.getTransaction(transaction.Id)

}

var transactionsData = Transactions{
	{Id: 1, Date: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC), AccountId: 1, Amount: 123.45, Payee: "Walmart", CategoryId: 1, Comment: ""},
	{Id: 2, Date: time.Date(2021, time.Month(3), 21, 1, 10, 30, 0, time.UTC), AccountId: 2, Amount: -150.45, Payee: "Walmart", CategoryId: 1, Comment: ""},
	{Id: 3, Date: time.Date(2021, time.Month(4), 21, 1, 10, 30, 0, time.UTC), AccountId: 3, Amount: 3.45, Payee: "Walmart", CategoryId: 1, Comment: ""},
	{Id: 4, Date: time.Date(2021, time.Month(5), 21, 1, 10, 30, 0, time.UTC), AccountId: 4, Amount: -3.45, Payee: "Walmart", CategoryId: 1, Comment: ""},
}
