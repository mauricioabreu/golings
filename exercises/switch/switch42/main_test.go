// switch42
// Make the tests pass!

// I AM NOT DONE
//
// hexVal returns the value of a hexadecimal digit, or -1.
// Practices a tagless switch with arithmetic on characters.
package main_test

import "testing"

func hexVal(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'a' && c <= 'f':
		return int(c - 'a')
	}
	return -1
}

func TestHexVal(t *testing.T) {
	cases := map[byte]int{'0': 0, '9': 9, 'a': 10, 'f': 15, 'B': 11, 'g': -1}
	for in, want := range cases {
		if got := hexVal(in); got != want {
			t.Errorf("hexVal(%c) = %d, want %d", in, got, want)
		}
	}
}
