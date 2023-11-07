package main

import "testing"

func TestWallet(t *testing.T) {
	given := Wallet{}

	given.Deposit(Bitcoin(10))

	when := given.Balance()
	then := Bitcoin(10)

	if when != then {
		t.Errorf("Given %#v when %d then %d", given, when, then)
	}
}
