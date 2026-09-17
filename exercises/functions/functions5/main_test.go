// functions5
// Make the tests pass!

// I AM NOT DONE
//
// withdraw must return an error with the text "insufficient funds".
// Practices building errors with errors.New.
package main_test

import (
	"errors"
	"testing"
)

func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, errors.New("not enough money")
	}
	return balance - amount, nil
}

func TestWithdraw(t *testing.T) {
	if b, err := withdraw(100, 30); err != nil || b != 70 {
		t.Errorf("withdraw(100, 30) = %d, %v", b, err)
	}
	_, err := withdraw(10, 30)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("withdraw(10, 30) error = %v, want insufficient funds", err)
	}
}
