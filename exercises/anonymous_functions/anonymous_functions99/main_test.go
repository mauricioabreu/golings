// anonymous_functions99
// Make the tests pass!

// I AM NOT DONE
//
// trimPunct strips the punctuation from both ends of a string.
// Practices a literal for strings.TrimFunc.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func trimPunct(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsPunct(r)
	})
}

func TestTrimPunct(t *testing.T) {
	if got := trimPunct("...hi, there!!"); got != "hi, there" {
		t.Errorf("trimPunct = %q", got)
	}
}
