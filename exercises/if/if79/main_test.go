// if79
// Make the tests pass!

// I AM NOT DONE
//
// shortage must return true when more is asked for than the warehouse holds.
// The condition can never hold: the difference of unsigned numbers is never negative.
// Practices comparing unsigned values.
package main_test

import "testing"

func shortage(stock, requested uint) bool {
	if stock-requested < 0 {
		return true
	}
	return false
}

func TestShortage(t *testing.T) {
	if !shortage(3, 5) {
		t.Errorf("shortage(3, 5) should be true")
	}
	if shortage(5, 5) || shortage(10, 1) {
		t.Errorf("no shortage expected")
	}
}
