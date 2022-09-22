package main

import "errors"

type WellAccount struct {
	balance int
}

func NewWellAccount() *WellAccount {
	return &WellAccount{
		balance: 0,
	}
}

func (w *WellAccount) GetBalance() int {
	return w.balance

}

func (w *WellAccount) Deposit(amount int) {
	w.balance += amount

}

func (w *WellAccount) Withdrawn(amount int) error {
	newAmount := w.balance - amount
	if newAmount < 0 {
		return errors.New("insufficient fund")

	}
	w.balance = newAmount
	return nil
}
