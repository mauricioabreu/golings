// variables89
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the price with tax. The code does not compile:
// the tax is computed but never used.
// Go rejects unused local variables.
package main_test

import "testing"

func withTax(price int) int {
	tax := price * 20 / 100
	return price
}

func TestWithTax(t *testing.T) {
	if got := withTax(100); got != 120 {
		t.Errorf("withTax(100) = %d, want 120", got)
	}
}
