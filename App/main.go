package main

import (
	transactions "App/transactions"
	"fmt"
)

func main() {
	trans := transactions.Transaction{AccountId: 9}
	trans.AccountId = 9
	//tranActionResult := transactions.Delete(1)
	if err := transactions.Update2(trans); err != nil {
		fmt.Printf("error. account id: %v\n", trans.AccountId)
	}

	fmt.Printf("success. account id: %v\n", trans.AccountId)
}
