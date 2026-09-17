// range72
// Make the tests pass!

// I AM NOT DONE
//
// countSpaces counts the spaces in a string.
// Practices comparing a rune in a range.
package main_test

import "testing"

func countSpaces(s string) int {
	n := 0
	for _, r := range s {
		if r == '_' {
			n++
		}
	}
	return n
}

func TestCountSpaces(t *testing.T) {
	if got := countSpaces("a b  c_d"); got != 3 {
		t.Errorf("countSpaces = %d, want 3", got)
	}
}
