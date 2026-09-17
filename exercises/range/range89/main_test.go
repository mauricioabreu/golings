// range89
// Make the tests pass!

// I AM NOT DONE
//
// reverse reverses a string by putting every rune in front of the result.
// Practices a range over the runes of a string.
package main_test

import "testing"

func reverse(s string) string {
	out := ""
	for _, r := range s {
		out = out + string(r)
	}
	return out
}

func TestReverse(t *testing.T) {
	cases := map[string]string{"abc": "cba", "φως": "ςωφ", "": ""}
	for in, want := range cases {
		if got := reverse(in); got != want {
			t.Errorf("reverse(%q) = %q, want %q", in, got, want)
		}
	}
}
