package pointers

import (
	"testing"
)

func TestWalletRef(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// func TestWallet(t *testing.T) {

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()
// 		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}
// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// }
