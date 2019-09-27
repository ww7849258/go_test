package bits

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amt
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
