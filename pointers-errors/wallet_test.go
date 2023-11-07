package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		given := Wallet{}

		given.Deposit(Bitcoin(10))

		when := given.Balance()
		then := Bitcoin(10)

		if when != then {
			t.Errorf("Given %#v when %s then %s", given, when, then)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		given := Wallet{balance: Bitcoin(20)}

		given.withdraw(Bitcoin(10))

		when := given.Balance()
		then := Bitcoin(10)

		if when != then {
			t.Errorf("Given %#v when %s then %s", given, when, then)
		}
	})
}
