package main

type Wallet struct {
	// lowercase means private
	balance int
}

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return 0
}
