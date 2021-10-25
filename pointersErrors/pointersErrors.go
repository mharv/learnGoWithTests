package main

import "fmt"

func main() {

	var A Wallet
	var B Wallet

	A.Deposit(100)
	B.Deposit(1100)

	fmt.Printf("Wallet A balance is equal to %d \n", A.Balance())
	fmt.Printf("Wallet B balance is equal to %d \n", B.Balance())

}

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
