package main

import "errors"

type BitcoinAccount struct {
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     300,
	}
}

func (b *BitcoinAccount) GetBalance() int {
	return b.balance

}

func (b *BitcoinAccount) Deposit(amount int) {
	b.balance += amount

}

func (b *BitcoinAccount) Withdrawn(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		errors.New("insufficient fund")

	}
	b.balance = newBalance
	return nil

}
