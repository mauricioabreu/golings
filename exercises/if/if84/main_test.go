// if84
// Make the tests pass!

// I AM NOT DONE
//
// sameTitle compares titles ignoring case and the spaces around them.
// Titles that are empty after trimming are never equal.
// Practices preparing the data before the condition.
package main_test

import (
	"strings"
	"testing"
)

func sameTitle(a, b string) bool {
	if strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}

func TestSameTitle(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{{"Go Tour", "  go tour ", true}, {"Go", "Rust", false}, {"  ", "", false}, {"", "", false}}
	for _, c := range cases {
		if got := sameTitle(c.a, c.b); got != c.want {
			t.Errorf("sameTitle(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
