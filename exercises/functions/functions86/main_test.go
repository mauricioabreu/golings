// functions86
// Make the tests pass!

// I AM NOT DONE
//
// newAccount returns a deposit and a balance function working on one account.
// Right now balance does not see the deposits.
// Closures share the variable itself, not a copy of it.
package main_test

import "testing"

func newAccount(initial int) (deposit func(int), balance func() int) {
	deposit = func(n int) { initial += n }
	snapshot := initial
	balance = func() int { return snapshot }
	return
}

func TestAccount(t *testing.T) {
	dep, bal := newAccount(100)
	dep(50)
	dep(25)
	if got := bal(); got != 175 {
		t.Errorf("balance = %d, want 175", got)
	}
}
