// range60
// Make the tests pass!

// I AM NOT DONE
//
// depositAll tops up every account by calling a method with a pointer receiver.
// The method is called on the loop variable, and the balances in the slice do not change.
// The v of a range is a copy, and a pointer method changes that copy.
package main_test

import "testing"

type Account struct{ Balance int }

func (a *Account) Deposit(n int) { a.Balance += n }

func depositAll(accounts []Account, n int) {
	for _, acc := range accounts {
		acc.Deposit(n)
	}
}

func TestDepositAll(t *testing.T) {
	accs := []Account{{10}, {20}}
	depositAll(accs, 5)
	if accs[0].Balance != 15 || accs[1].Balance != 25 {
		t.Errorf("balances = %v", accs)
	}
}
