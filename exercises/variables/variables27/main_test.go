// variables27
// Make the tests pass!

// I AM NOT DONE
//
// bitInfo must return the number of set bits and the length of the number in bits.
// Practices the math/bits package and unsigned types.
package main_test

import (
	"math/bits"
	"testing"
)

func bitInfo(v uint) (ones, length int) {
	ones = bits.TrailingZeros(v)
	length = bits.LeadingZeros(v)
	return
}

func TestBitInfo(t *testing.T) {
	cases := []struct {
		v            uint
		ones, length int
	}{{0b1011, 3, 4}, {1, 1, 1}, {0, 0, 0}, {255, 8, 8}}
	for _, c := range cases {
		ones, length := bitInfo(c.v)
		if ones != c.ones || length != c.length {
			t.Errorf("bitInfo(%b) = %d, %d; want %d, %d", c.v, ones, length, c.ones, c.length)
		}
	}
}
