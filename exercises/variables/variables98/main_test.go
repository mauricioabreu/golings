// variables98
// Make the tests pass!

// I AM NOT DONE
//
// lowNibble must return the low 4 bits of a number.
// The mask is written as a binary literal, and it is wrong.
// Practices 0b binary literals and the bitwise AND.
package main_test

import "testing"

func lowNibble(v uint8) uint8 {
	const mask = 0b1010
	return v & mask
}

func TestLowNibble(t *testing.T) {
	cases := map[uint8]uint8{0xAB: 0xB, 0x0F: 0xF, 0xF0: 0, 0x37: 7}
	for in, want := range cases {
		if got := lowNibble(in); got != want {
			t.Errorf("lowNibble(%#x) = %#x, want %#x", in, got, want)
		}
	}
}
