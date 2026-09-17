// range81
// Make the tests pass!

// I AM NOT DONE
//
// onlyLetters collects the letters of a string into a []byte. For non-ASCII letters
// it produces junk.
// byte(r) truncates a rune to a single byte.
package main_test

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func onlyLetters(s string) []byte {
	var b []byte
	for _, r := range s {
		if unicode.IsLetter(r) {
			b = append(b, byte(r))
		}
	}
	return b
}

func TestOnlyLetters(t *testing.T) {
	_ = utf8.RuneLen
	if got := string(onlyLetters("Go-1 λ Ω!")); got != "GoλΩ" {
		t.Errorf("onlyLetters = %q, want GoλΩ", got)
	}
}
