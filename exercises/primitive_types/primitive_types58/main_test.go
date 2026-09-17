// primitive_types58
// Make the tests pass!

// I AM NOT DONE
//
// popcount must count the set bits in every byte of a slice without math/bits.
// Practices bitwise operations on a byte.
package main_test

import "testing"

func popcount(data []byte) int {
	n := 0
	for _, b := range data {
		if b&1 == 1 {
			n++
		}
	}
	return n
}

func TestPopcount(t *testing.T) {
	cases := []struct {
		in   []byte
		want int
	}{{[]byte{0xFF}, 8}, {[]byte{1, 2, 3}, 4}, {nil, 0}, {[]byte{0x80, 0x01}, 2}}
	for _, c := range cases {
		if got := popcount(c.in); got != c.want {
			t.Errorf("popcount(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
