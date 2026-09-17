// primitive_types82
// Make the tests pass!

// I AM NOT DONE
//
// stripPrefix removes the prefix "abc-" from an identifier exactly once.
// For "abc-cab-1" it currently removes too much.
// TrimLeft takes a set of characters (a cutset), not a substring.
package main_test

import (
	"strings"
	"testing"
)

func stripPrefix(id string) string {
	return strings.TrimLeft(id, "abc-")
}

func TestStripPrefix(t *testing.T) {
	cases := map[string]string{"abc-cab-1": "cab-1", "abc-42": "42", "bca-7": "bca-7"}
	for in, want := range cases {
		if got := stripPrefix(in); got != want {
			t.Errorf("stripPrefix(%q) = %q, want %q", in, got, want)
		}
	}
}
