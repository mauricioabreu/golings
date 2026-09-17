// primitive_types86
// Make the tests pass!

// I AM NOT DONE
//
// countA must count how many times the letter 'a' occurs in a string.
// Practices strings.Count.
package main_test

import (
	"strings"
	"testing"
)

func countA(s string) int {
	return strings.Count(s, "A")
}

func TestCountA(t *testing.T) {
	cases := map[string]int{"banana": 3, "Apple": 0, "": 0}
	for in, want := range cases {
		if got := countA(in); got != want {
			t.Errorf("countA(%q) = %d, want %d", in, got, want)
		}
	}
}
