// if11
// Make the tests pass!

// I AM NOT DONE
//
// withdraw takes the amount out when the balance plus the overdraft limit allows it.
// A negative amount is the error "invalid amount".
// Practices checking several conditions with different errors.
package main_test

import (
	"errors"
	"testing"
)

func withdraw(balance, limit, amount int) (int, error) {
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	return balance - amount, nil
}

func TestWithdraw(t *testing.T) {
	if b, err := withdraw(100, 50, 140); err != nil || b != -40 {
		t.Errorf("withdraw(100, 50, 140) = %d, %v; want -40", b, err)
	}
	if _, err := withdraw(100, 50, 151); err == nil || err.Error() != "insufficient funds" {
		t.Errorf("withdraw over limit err = %v", err)
	}
	if _, err := withdraw(100, 50, -5); err == nil || err.Error() != "invalid amount" {
		t.Errorf("withdraw(-5) err = %v", err)
	}
}
