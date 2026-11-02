// closures3
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// newAccount returns two functions that share one balance: deposit adds to it
// and balance reads it.
func newAccount() (deposit func(int), balance func() int) {
	total := 0

	deposit = func(amount int) {
		// TODO: Add the amount to the total that balance reads.
		total := total + amount
		_ = total
	}
	balance = func() int {
		return total
	}

	return deposit, balance
}

func TestAccount(t *testing.T) {
	deposit, balance := newAccount()

	deposit(10)
	deposit(5)

	if got := balance(); got != 15 {
		t.Errorf("balance = %d, want 15", got)
	}
}

func TestAccountsDoNotShareTheirBalance(t *testing.T) {
	depositA, balanceA := newAccount()
	_, balanceB := newAccount()

	depositA(100)

	if got := balanceB(); got != 0 {
		t.Errorf("the other account has %d, want 0", got)
	}
	if got := balanceA(); got != 100 {
		t.Errorf("balance = %d, want 100", got)
	}
}
