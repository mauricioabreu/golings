// if67
// Make the tests pass!

// I AM NOT DONE
//
// strong must require a length of 8 or more, at least one digit and at least one capital letter.
// Practices flags computed in a loop and a final condition.
package main_test

import (
	"testing"
	"unicode"
)

func strong(pw string) bool {
	hasDigit, hasUpper := false, false
	for _, r := range pw {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return len(pw) >= 8 || hasDigit
}

func TestStrong(t *testing.T) {
	cases := map[string]bool{"Passw0rd": true, "password": false, "Password": false, "passw0rd": false, "P4ss": false}
	for in, want := range cases {
		if got := strong(in); got != want {
			t.Errorf("strong(%q) = %v, want %v", in, got, want)
		}
	}
}
