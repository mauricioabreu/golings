// primitive_types78
// Make the tests pass!

// I AM NOT DONE
//
// swapCase flips the case of every letter.
// Practices unicode.IsUpper, ToLower and ToUpper.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func swapCase(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, s)
}

func TestSwapCase(t *testing.T) {
	cases := map[string]string{"Hello World": "hELLO wORLD", "ΚαΛά": "κΑλΆ", "123": "123"}
	for in, want := range cases {
		if got := swapCase(in); got != want {
			t.Errorf("swapCase(%q) = %q, want %q", in, got, want)
		}
	}
}
