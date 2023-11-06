package main

import "testing"

func TestWallet(t *testing.T) {
	given := Wallet{}

	given.Deposit(10)

	when := given.Balance()
	then := 10

	if when != then {
		t.Errorf("Given %#v when %d then %d", given, when, then)
	}
}
