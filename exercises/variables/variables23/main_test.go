// variables23
// Make the tests pass!

// I AM NOT DONE
//
// This function must return how much stock is left, but never less than zero.
// When the order is larger than the stock the result is a huge number.
// Practices unsigned overflow on subtraction.
package main_test

import "testing"

func remaining(stock, order uint) uint {
	left := stock - order
	if left < 0 {
		return 0
	}
	return left
}

func TestRemaining(t *testing.T) {
	if got := remaining(10, 3); got != 7 {
		t.Errorf("remaining(10, 3) = %d, want 7", got)
	}
	if got := remaining(3, 10); got != 0 {
		t.Errorf("remaining(3, 10) = %d, want 0", got)
	}
}
