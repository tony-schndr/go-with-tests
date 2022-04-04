package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got != insuffuscientFundsError {
			t.Fatal("expected overdraft error")
		}
		if got.Error() != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw with insuffuscient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		err := wallet.Withdraw(10)
		assertError(t, err, insuffuscientFundsError.Error())
		assertBalance(t, wallet, Bitcoin(5))
	})
}
