// variables39
// Make the tests pass!

// I AM NOT DONE
//
// The variable op must hold the multiplication function.
// It is left uninitialized, so calling it panics.
// Practices functions as values and the zero value of a func variable.
package main_test

import "testing"

func multiply(a, b int) int { return a * b }

func apply(a, b int) int {
	var op func(int, int) int
	return op(a, b)
}

func TestApply(t *testing.T) {
	if got := apply(6, 7); got != 42 {
		t.Errorf("apply(6, 7) = %d, want 42", got)
	}
}
