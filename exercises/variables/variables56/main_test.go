// variables56
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the first letter of the alphabet as a rune.
// Practices the difference between a rune literal and a string literal.
package main_test

import "testing"

func firstLetter() rune {
	var letter rune = "A"
	return letter
}

func TestFirstLetter(t *testing.T) {
	if got := firstLetter(); got != 65 {
		t.Errorf("firstLetter() = %d, want 65", got)
	}
}
