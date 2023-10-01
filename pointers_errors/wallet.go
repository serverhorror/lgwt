package pointers_errors

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("you don't have enough money")
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.02f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return fmt.Errorf("cannot withdraw: %w", ErrInsufficientFunds)
	}
	return w.Deposit(amount * -1)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
