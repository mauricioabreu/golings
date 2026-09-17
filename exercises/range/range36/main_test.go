// range36
// Make the tests pass!

// I AM NOT DONE
//
// shiftGreek moves lowercase greek letters k places around the alphabet.
// Practices a range over runes and rune arithmetic.
package main_test

import (
	"strings"
	"testing"
)

func shiftGreek(s string, k int) string {
	var b strings.Builder
	for _, r := range s {
		if r >= 'α' && r < 'ω' {
			r = r + rune(k)
		}
		b.WriteRune(r)
	}
	return b.String()
}

func TestShiftGreek(t *testing.T) {
	cases := map[string]string{"αβγ": "βγδ", "ωμα!": "ανβ!"}
	for in, want := range cases {
		if got := shiftGreek(in, 1); got != want {
			t.Errorf("shiftGreek(%q) = %q, want %q", in, got, want)
		}
	}
}
