// primitive_types81
// Make the tests pass!

// I AM NOT DONE
//
// bits8 must return the binary form of a byte as exactly 8 digits.
// Practices width and zero padding in fmt.
package main_test

import (
	"fmt"
	"testing"
)

func bits8(b byte) string {
	return fmt.Sprintf("%8b", b)
}

func TestBits8(t *testing.T) {
	cases := map[byte]string{5: "00000101", 255: "11111111", 0: "00000000"}
	for in, want := range cases {
		if got := bits8(in); got != want {
			t.Errorf("bits8(%d) = %q, want %q", in, got, want)
		}
	}
}
