package transactions

import (
	"encoding/json"
	"io"
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	Date       time.Time `json:"date"`
	AccountId  int       `json:"accountId"`
	Amount     float32   `json:"amount"`
	Payee      string    `json:"payee"`
	CategoryId int       `json:"category"`
	Comment    string    `json:"comment"`
}

func (t *Transaction) toJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(t)
}

func (t *Transaction) fromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(t)
}

type Transactions []*Transaction

func (ts *Transactions) toJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(ts)
}

func (ts *Transactions) fromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(ts)
}
