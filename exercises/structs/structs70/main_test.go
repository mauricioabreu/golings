// structs70
// Make the tests pass!

// I AM NOT DONE
//
// Account.Balance returns the balance in whole units, from cents.
// Practices getter methods and unexported fields.
package main_test

import "testing"

type Account struct{ cents int }

func (a Account) Balance() int {
	return a.cents * 100
}

func TestBalance(t *testing.T) {
	if got := (Account{cents: 12345}).Balance(); got != 123 {
		t.Errorf("Balance = %d", got)
	}
}
