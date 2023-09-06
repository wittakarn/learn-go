package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		actual := wallet.Balance()

		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("actual %d expected %d", actual, expected)
		}
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		actual := wallet.Balance()

		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("actual %d expected %d", actual, expected)
		}
	})
}
