// range49
// Make the tests pass!

// I AM NOT DONE
//
// countDigits counts the digits in a string.
// Practices a range over a string and unicode.IsDigit.
package main_test

import (
	"testing"
	"unicode"
)

func countDigits(s string) int {
	n := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			n++
		}
	}
	return n
}

func TestCountDigits(t *testing.T) {
	if got := countDigits("room 42, floor 7"); got != 3 {
		t.Errorf("countDigits = %d, want 3", got)
	}
}
