package main

import "fmt"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdrawn(amount int) error
}

func main() {
	myAccounts := []IBankAccount{
		NewBitcoinAccount(),
		NewWellAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withdrawn(80); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}

	}

	for _, account := range myAccounts {
		balance := account.GetBalance()
		fmt.Printf("balance = %d\n", balance)

	}

}
