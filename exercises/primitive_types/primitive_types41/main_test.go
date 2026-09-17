// primitive_types41
// Make the tests pass!

// I AM NOT DONE
//
// bitsOf returns the 8 bits of an int8 in two's complement: -1 -> "11111111".
// Practices converting an int8 to a uint8 and formatting it.
package main_test

import (
	"fmt"
	"testing"
)

func bitsOf(v int8) string {
	u := v
	return fmt.Sprintf("%b", u)
}

func TestBitsOf(t *testing.T) {
	cases := map[int8]string{-1: "11111111", 5: "00000101", -128: "10000000", 127: "01111111"}
	for in, want := range cases {
		if got := bitsOf(in); got != want {
			t.Errorf("bitsOf(%d) = %s, want %s", in, got, want)
		}
	}
}
