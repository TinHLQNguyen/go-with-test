package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, given Wallet, then Bitcoin) {
		t.Helper()
		when := given.Balance()
		if when != then {
			t.Errorf("Given %#v when %s then %s", given, when, then)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		given := Wallet{}

		given.Deposit(Bitcoin(10))
		then := Bitcoin(10)

		assertBalance(t, given, then)
	})
	t.Run("withdraw", func(t *testing.T) {
		given := Wallet{balance: Bitcoin(20)}

		given.Withdraw(Bitcoin(10))
		then := Bitcoin(10)

		assertBalance(t, given, then)
	})
}
