// anonymous_functions56
// Make the tests pass!

// I AM NOT DONE
//
// adder returns a closure adding up the values it is given.
// Practices a closure with mutable state.
package main_test

import "testing"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum = x
		return sum
	}
}

func TestAdder(t *testing.T) {
	a := adder()
	a(10)
	a(5)
	if got := a(1); got != 16 {
		t.Errorf("sum = %d, want 16", got)
	}
}
