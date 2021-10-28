package main

import (
	"errors"
	"fmt"
)

func main() {

	var A Wallet
	A.Deposit(10)
	A.Withdraw(11)
	fmt.Println(A.balance.String())

}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("uh oh")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
