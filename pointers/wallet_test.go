package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if err != want {
			t.Errorf("got %q but want %q", err, want)
		}
	}
	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("got an error but didn't want one: %q", err)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
