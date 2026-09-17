// if9
// Make the tests pass!

// I AM NOT DONE
//
// validEmail: exactly one @, not in the first position, with a dot after it
// that is neither right after the @ nor at the very end.
// Practices a sequence of checks with early returns.
package main_test

import (
	"strings"
	"testing"
)

func validEmail(s string) bool {
	at := strings.Index(s, "@")
	if at < 0 {
		return false
	}
	return strings.Contains(s, ".")
}

func TestValidEmail(t *testing.T) {
	cases := map[string]bool{
		"ann@mail.com": true, "a@b.co": true, "@mail.com": false, "ann.mail.com": false,
		"ann@mail": false, "ann@.com": false, "ann@mail.": false, "a@b@c.com": false,
	}
	for in, want := range cases {
		if got := validEmail(in); got != want {
			t.Errorf("validEmail(%q) = %v, want %v", in, got, want)
		}
	}
}
