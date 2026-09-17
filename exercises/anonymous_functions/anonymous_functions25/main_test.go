// anonymous_functions25
// Make the tests pass!

// I AM NOT DONE
//
// compute has to return the value raised by a deferred literal.
// The literal changes a local variable, and the result has already been computed.
// A return with an expression fixes the value before the defer runs.
package main_test

import "testing"

func compute(x int) int {
	result := x * 2
	defer func() { result++ }()
	return result
}

func TestCompute(t *testing.T) {
	if got := compute(5); got != 11 {
		t.Errorf("compute(5) = %d, want 11", got)
	}
}
