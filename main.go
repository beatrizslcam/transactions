package main

import (
	"fmt"
	"transactions/domain/account"
	"transactions/domain/transfers"
)

func main() {
	test := transfers.CreateTransfer("1", "2", 100)
	fmt.Print(test)

	newaccount := account.CreateAccount("josé", "123456789")
	fmt.Print(newaccount)

}