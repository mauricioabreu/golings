// functions84
// Make the tests pass!

// I AM NOT DONE
//
// makeMultiplier(n) must return a function that multiplies its argument by n.
// Practices functions returning functions.
package main_test

import "testing"

func makeMultiplier(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func TestMakeMultiplier(t *testing.T) {
	triple := makeMultiplier(3)
	if got := triple(5); got != 15 {
		t.Errorf("triple(5) = %d, want 15", got)
	}
	if got := makeMultiplier(0)(7); got != 0 {
		t.Errorf("makeMultiplier(0)(7) = %d, want 0", got)
	}
}
