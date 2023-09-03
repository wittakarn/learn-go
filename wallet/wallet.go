package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amont int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance = w.balance + amont
}

func (w *Wallet) Balance() int {
	return w.balance
}
