package wallet

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amont Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance = w.balance + amont
}

func (w *Wallet) Withdraw(amont Bitcoin) {
	fmt.Printf("address of balance in Withdraw is %v \n", &w.balance)
	w.balance = w.balance - amont
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
