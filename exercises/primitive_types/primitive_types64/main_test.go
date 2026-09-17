// primitive_types64
// Make the tests pass!

// I AM NOT DONE
//
// isPalindrome checks a phrase, ignoring case and every character but letters.
// It works for any alphabet.
// Practices []rune, unicode.IsLetter and unicode.ToLower.
package main_test

import (
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		letters = append(letters, r)
	}
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		if letters[i] != letters[j] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	_ = unicode.IsLetter
	cases := map[string]bool{"A man, a plan, a canal: Panama": true, "No 'x' in Nixon": true, "Go": false, "": true}
	for in, want := range cases {
		if got := isPalindrome(in); got != want {
			t.Errorf("isPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}
