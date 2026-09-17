// primitive_types25
// Make the tests pass!

// I AM NOT DONE
//
// letters counts the letters of any alphabet, skipping digits, spaces and punctuation.
// Practices iterating over runes and unicode.IsLetter.
package main_test

import (
	"testing"
	"unicode"
)

func letters(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			n++
		}
	}
	return n
}

func TestLetters(t *testing.T) {
	_ = unicode.IsLetter
	cases := map[string]int{"Go 1.22!": 2, "Γειά, κόσμε": 9, "": 0}
	for in, want := range cases {
		if got := letters(in); got != want {
			t.Errorf("letters(%q) = %d, want %d", in, got, want)
		}
	}
}
