// variables83
// Make the tests pass!

// I AM NOT DONE
//
// shift must rotate a lowercase latin letter k positions around the alphabet.
// Right now letters near the end of the alphabet turn into stray characters.
// Practices byte arithmetic and the remainder operator.
package main_test

import "testing"

func shift(c byte, k int) byte {
	return c + byte(k)
}

func TestShift(t *testing.T) {
	cases := []struct {
		c    byte
		k    int
		want byte
	}{{'a', 1, 'b'}, {'x', 3, 'a'}, {'z', 27, 'a'}, {'m', 0, 'm'}}
	for _, c := range cases {
		if got := shift(c.c, c.k); got != c.want {
			t.Errorf("shift(%c, %d) = %c, want %c", c.c, c.k, got, c.want)
		}
	}
}
