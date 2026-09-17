// anonymous_functions75
// Make the tests pass!

// I AM NOT DONE
//
// noDigits removes the digits from a string with strings.Map.
// Practices a literal returning -1 to drop a character.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func noDigits(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return ' '
		}
		return r
	}, s)
}

func TestNoDigits(t *testing.T) {
	if got := noDigits("a1b22c"); got != "abc" {
		t.Errorf("noDigits = %q", got)
	}
}
