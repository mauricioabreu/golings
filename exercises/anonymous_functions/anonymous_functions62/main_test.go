// anonymous_functions62
// Make the tests pass!

// I AM NOT DONE
//
// lastDigit returns the index of the last digit through strings.LastIndexFunc.
// Practices a literal for the search-from-the-end functions.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func lastDigit(s string) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		return unicode.IsLetter(r)
	})
}

func TestLastDigit(t *testing.T) {
	if lastDigit("a1b2c") != 3 || lastDigit("abc") != -1 {
		t.Errorf("lastDigit works incorrectly")
	}
}
