package pointers

import "testing"

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Expected %q and got %q", want, got)
	}
}

func assertNoError(t testing.TB, got error) {

	t.Helper()

	if got != nil {
		t.Fatal("Unexpected error")
	}

}

func assertError(t testing.TB, got error, want error) {

	t.Helper()

	if got == nil {
		t.Fatal("Expected error but didn't get one")
	}

	if got != want {
		t.Errorf("Expected %q and got %q", want, got)
	}

}

func TestWallet(t *testing.T) {

	t.Run("Deposit()", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw()", func(t *testing.T) {

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)

	})

	t.Run("Withdraw() error", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)

	})

}

