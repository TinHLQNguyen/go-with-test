package main

import "testing"

func TestWallet(t *testing.T) {
	given := Wallet{}

	given.Deposit(Bitcoin(10))

	when := given.Balance()
	then := Bitcoin(10)

	if when != then {
		t.Errorf("Given %#v when %s then %s", given, when, then)
	}
}
