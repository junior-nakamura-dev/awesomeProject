package main

import "fmt"

type Account struct {
	credit int
}

func addCredit(account *Account, value int) {
	account.credit += value
}

func main() {
	account := Account{credit: 100}

	fmt.Println("Account: ", account)

	addCredit(&account, 50)
	fmt.Println("Account after add credit: ", account)
}
