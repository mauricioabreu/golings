// if3
// Make the tests pass!

// I AM NOT DONE
//
// isUpper must return true for uppercase latin letters.
// Practices comparing runes against a range of characters.
package main_test

import "testing"

func isUpper(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func TestIsUpper(t *testing.T) {
	cases := map[rune]bool{'A': true, 'Z': true, 'a': false, '1': false, 'm': false}
	for in, want := range cases {
		if got := isUpper(in); got != want {
			t.Errorf("isUpper(%q) = %v, want %v", in, got, want)
		}
	}
}
