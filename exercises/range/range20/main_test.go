// range20
// Make the tests pass!

// I AM NOT DONE
//
// indexOfE returns the byte position of the letter 'é' in a string, or -1.
// The index of a range over a string is a byte offset.
package main_test

import "testing"

func indexOfE(s string) int {
	pos := 0
	for i, r := range s {
		if r == 'é' {
			return pos
		}
		pos++
	}
	return -1
}

func TestIndexOfE(t *testing.T) {
	cases := map[string]int{"café": 3, "ééé": 0, "cafe": -1, "ñé": 2}
	for in, want := range cases {
		if got := indexOfE(in); got != want {
			t.Errorf("indexOfE(%q) = %d, want %d", in, got, want)
		}
	}
}
