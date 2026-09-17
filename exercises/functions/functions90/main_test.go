// functions90
// Make the tests pass!

// I AM NOT DONE
//
// firstDigit must return the index of the first digit in a string, or -1.
// Practices strings.IndexFunc and the functions of the unicode package.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func firstDigit(s string) int {
	return strings.IndexFunc(s, unicode.IsLetter)
}

func TestFirstDigit(t *testing.T) {
	cases := map[string]int{"abc1": 3, "7up": 0, "none": -1, "": -1}
	for in, want := range cases {
		if got := firstDigit(in); got != want {
			t.Errorf("firstDigit(%q) = %d, want %d", in, got, want)
		}
	}
}
