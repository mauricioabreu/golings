// primitive_types85
// Make the tests pass!

// I AM NOT DONE
//
// highBitSet checks whether the top bit of an int8 is set.
// The code does not compile: the mask 0x80 does not fit in an int8.
// Practices typed operands and the range of the constants.
package main_test

import "testing"

func highBitSet(v int8) bool {
	return v&0x80 != 0
}

func TestHighBitSet(t *testing.T) {
	cases := map[int8]bool{-1: true, -128: true, 127: false, 0: false, 64: false}
	for in, want := range cases {
		if got := highBitSet(in); got != want {
			t.Errorf("highBitSet(%d) = %v, want %v", in, got, want)
		}
	}
}
