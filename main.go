package main

import (
	"fmt"

	"github.com/NoTabaco/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kimchi")
	fmt.Println(account)
}
