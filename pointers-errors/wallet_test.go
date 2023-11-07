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
	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("wanted an error but didn't get")
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
	t.Run("withdraw more than fund", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		given := Wallet{balance: startingBalance}

		err := given.Withdraw(Bitcoin(100))
		then := startingBalance

		assertError(t, err)
		assertBalance(t, given, then)
	})
}
