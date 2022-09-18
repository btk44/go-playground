package transactions

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	Date       time.Time `json:"date"`
	AccountId  int       `json:"accountId"`
	Amount     float32   `json:"amount"`
	Payee      string    `json:"payee"`
	CategoryId int       `json:"category"`
	Comment    string    `json:"comment"`
}
