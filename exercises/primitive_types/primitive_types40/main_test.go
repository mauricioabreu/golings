// primitive_types40
// Make the tests pass!

// I AM NOT DONE
//
// alphabetSize must return the number of letters from 'a' to 'z' inclusive.
// Practices arithmetic on rune constants.
package main_test

import "testing"

func alphabetSize() int {
	return int('z' - 'a')
}

func TestAlphabetSize(t *testing.T) {
	if got := alphabetSize(); got != 26 {
		t.Errorf("alphabetSize() = %d, want 26", got)
	}
}
