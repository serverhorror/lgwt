package pointers_errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestWalletString(t *testing.T) {
	wallet := Wallet{}
	if err := wallet.Deposit(10); err != nil {
		t.Error(err)
	}
	got := fmt.Sprint(wallet.Balance())
	want := "10.00 BTC"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWalletOperations(t *testing.T) {
	type operation func(w *Wallet, amount Bitcoin) error

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	// Use functions to redirect to the right method
	deposit := func(w *Wallet, amount Bitcoin) error {
		return w.Deposit(amount)

	}

	withdraw := func(w *Wallet, amount Bitcoin) error {
		return w.Withdraw(amount)
	}

	testCases := []struct {
		name    string
		wallet  Wallet
		amount  Bitcoin
		op      operation
		want    Bitcoin
		wantErr error
	}{
		{"Withdraw with funds", Wallet{balance: Bitcoin(20)}, Bitcoin(10), withdraw, Bitcoin(10), nil},
		{"Withdraw without funds", Wallet{balance: Bitcoin(20)}, Bitcoin(30), withdraw, Bitcoin(20), ErrInsufficientFunds},
		{"Deposit", Wallet{balance: Bitcoin(20)}, Bitcoin(10), deposit, Bitcoin(30), nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.op(&tc.wallet, tc.amount); !errors.Is(err, tc.wantErr) {
				t.Error(err)
			}
			assertBalance(t, tc.wallet, tc.want)
		})
	}
}
