// primitive_types59
// Make the tests pass!

// I AM NOT DONE
//
// startsUpper must return true when the first rune is a capital letter of any alphabet.
// Practices the unicode package.
package main_test

import (
	"testing"
	"unicode"
)

func startsUpper(s string) bool {
	for _, r := range s {
		return unicode.IsLower(r)
	}
	return false
}

func TestStartsUpper(t *testing.T) {
	cases := map[string]bool{"Go": true, "Ωμέγα": true, "go": false, "1st": false, "": false}
	for in, want := range cases {
		if got := startsUpper(in); got != want {
			t.Errorf("startsUpper(%q) = %v, want %v", in, got, want)
		}
	}
}
