// variables36
// Make the tests pass!

// I AM NOT DONE
//
// The function returns a quotient and a remainder, and only the remainder is wanted.
// The wrong value ends up in the variable.
// Practices the blank identifier _ in a multiple assignment.
package main_test

import "testing"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func remainder(a, b int) int {
	r, _ := divmod(a, b)
	return r
}

func TestRemainder(t *testing.T) {
	if got := remainder(17, 5); got != 2 {
		t.Errorf("remainder(17, 5) = %d, want 2", got)
	}
	if got := remainder(9, 3); got != 0 {
		t.Errorf("remainder(9, 3) = %d, want 0", got)
	}
}
