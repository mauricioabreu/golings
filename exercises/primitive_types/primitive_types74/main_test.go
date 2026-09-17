// primitive_types74
// Make the tests pass!

// I AM NOT DONE
//
// nextLetter must return the letter following the given one ('a' -> 'b').
// Practices byte arithmetic.
package main_test

import "testing"

func nextLetter(c byte) byte {
	return c - 1
}

func TestNextLetter(t *testing.T) {
	if got := nextLetter('a'); got != 'b' {
		t.Errorf("nextLetter('a') = %c, want b", got)
	}
	if got := nextLetter('Y'); got != 'Z' {
		t.Errorf("nextLetter('Y') = %c, want Z", got)
	}
}
