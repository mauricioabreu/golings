// primitive_types31
// Make the tests pass!

// I AM NOT DONE
//
// wordCount must count words separated by any number of spaces.
// Practices strings.Fields.
package main_test

import (
	"strings"
	"testing"
)

func wordCount(s string) int {
	return len(strings.Split(s, " "))
}

func TestWordCount(t *testing.T) {
	cases := map[string]int{"go is fun": 3, "  spaced   out  ": 2, "": 0}
	for in, want := range cases {
		if got := wordCount(in); got != want {
			t.Errorf("wordCount(%q) = %d, want %d", in, got, want)
		}
	}
}
