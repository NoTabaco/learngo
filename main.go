package main

import (
	"fmt"

	"github.com/NoTabaco/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kimchi")
	account.Deposit(1000)
	fmt.Println(account)
}
