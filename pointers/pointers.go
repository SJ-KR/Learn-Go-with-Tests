package pointers

import "fmt"

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
