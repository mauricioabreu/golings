// anonymous_functions42
// Make the tests pass!

// I AM NOT DONE
//
// Account.Depositor returns a deposit literal. The deposits never show on the account:
// the method has a value receiver and the literal captured a copy.
// A closure captures the receiver variable.
package main_test

import "testing"

type Account struct{ Balance int }

func (a Account) Depositor() func(int) {
	return func(n int) { a.Balance += n }
}

func TestDepositor(t *testing.T) {
	acc := &Account{}
	dep := acc.Depositor()
	dep(10)
	dep(5)
	if acc.Balance != 15 {
		t.Errorf("Balance = %d, want 15", acc.Balance)
	}
}
