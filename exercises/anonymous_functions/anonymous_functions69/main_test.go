// anonymous_functions69
// Make the tests pass!

// I AM NOT DONE
//
// firstUpper returns the index of the first capital letter through strings.IndexFunc.
// Practices a predicate literal.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func firstUpper(s string) int {
	return strings.IndexFunc(s, func(r rune) bool {
		return unicode.IsLower(r)
	})
}

func TestFirstUpper(t *testing.T) {
	if firstUpper("helloWorld") != 5 || firstUpper("abc") != -1 {
		t.Errorf("firstUpper works incorrectly")
	}
}
