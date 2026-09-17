// functions73
// Make the tests pass!

// I AM NOT DONE
//
// rot13 must encode latin letters and leave every other character alone.
// A mapping function is handed to strings.Map, and it is the wrong one.
// Practices higher order functions from the standard library.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func rotRune(r rune) rune {
	if !unicode.IsLetter(r) {
		return r
	}
	switch {
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	}
	return r
}

func rot13(s string) string {
	return strings.Map(unicode.ToUpper, s)
}

func TestRot13(t *testing.T) {
	if got := rot13("Hello, Go!"); got != "Uryyb, Tb!" {
		t.Errorf("rot13(Hello, Go!) = %q, want %q", got, "Uryyb, Tb!")
	}
}
