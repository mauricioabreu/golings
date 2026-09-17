// structs58
// Make the tests pass!

// I AM NOT DONE
//
// Transfer moves money between accounts; with too little money it returns an error
// and leaves the balances alone.
// Practices methods changing several structs.
package main_test

import (
	"errors"
	"testing"
)

type Account struct{ Balance int }

func Transfer(from, to *Account, amount int) error {
	from.Balance -= amount
	return nil
}

func TestTransfer(t *testing.T) {
	_ = errors.New
	a, b := &Account{100}, &Account{0}
	if err := Transfer(a, b, 30); err != nil || a.Balance != 70 || b.Balance != 30 {
		t.Errorf("after transfer: %v, %d, %d", err, a.Balance, b.Balance)
	}
	if err := Transfer(a, b, 500); err == nil || a.Balance != 70 || b.Balance != 30 {
		t.Errorf("failed transfer changed balances: %v, %d, %d", err, a.Balance, b.Balance)
	}
}
