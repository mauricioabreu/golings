// variables19
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the starting balance of a new account, which is 100.
// The variable is declared with var and no initializer, so it holds the zero value.
package main_test

import "testing"

func startBalance() int {
	var balance int
	return balance
}

func TestStartBalance(t *testing.T) {
	if got := startBalance(); got != 100 {
		t.Errorf("startBalance() = %d, want 100", got)
	}
}
