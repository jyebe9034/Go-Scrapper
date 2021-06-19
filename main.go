package main

import (
	"fmt"

	"github.com/jyebe/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hannah")
	fmt.Println(account)
}
