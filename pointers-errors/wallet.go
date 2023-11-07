package main

// create type from generic one
type Bitcoin int

type Wallet struct {
	// lowercase means private
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
