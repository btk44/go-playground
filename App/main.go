package main

import (
	transactions "App/transactions"
	"fmt"
)

func main() {
	tranActionResult := transactions.Add(*new(transactions.Transaction))
	//tranActionResult := transactions.Delete(1)

	if tranActionResult.IsSuccess() {
		fmt.Println("done")
	} else {
		fmt.Println(tranActionResult.ErrorMessage)
	}

}
