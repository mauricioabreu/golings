// range39
// Make the tests pass!

// I AM NOT DONE
//
// isPalindrome checks a string ignoring case, comparing runes from both ends.
// Practices a range over a []rune with an index.
package main_test

import (
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	rs := []rune(s)
	for i, r := range rs {
		if r != rs[len(rs)-i] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	_ = unicode.ToLower
	cases := map[string]bool{"Αββα": true, "Level": true, "go": false, "": true}
	for in, want := range cases {
		if got := isPalindrome(in); got != want {
			t.Errorf("isPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}
