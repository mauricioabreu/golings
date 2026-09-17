// range12
// Make the tests pass!

// I AM NOT DONE
//
// firstSpace returns the byte position of the first space, or -1.
// The index of a range over a string is a byte position.
package main_test

import "testing"

func firstSpace(s string) int {
	for i, r := range s {
		if r == ' ' {
			continue
		}
	}
	return -1
}

func TestFirstSpace(t *testing.T) {
	cases := map[string]int{"hello world": 5, "go": -1, " x": 0}
	for in, want := range cases {
		if got := firstSpace(in); got != want {
			t.Errorf("firstSpace(%q) = %d, want %d", in, got, want)
		}
	}
}
