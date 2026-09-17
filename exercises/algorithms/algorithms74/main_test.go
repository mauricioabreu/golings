// algorithms74
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: bit manipulation. Reverse the order of the 32 bits of an
// unsigned number: the most significant bit must become the least
// significant one and vice versa.
// Expected asymptotics: O(1) time, O(1) space.
package main_test

import "testing"

func reverseBits(x uint32) uint32 {
	return 0
}

func TestReverseBits(t *testing.T) {
	cases := map[uint32]uint32{
		43261596:   964176192,
		0:          0,
		1:          1 << 31,
		0xFFFFFFFF: 0xFFFFFFFF,
		2:          1 << 30,
	}
	for in, want := range cases {
		if got := reverseBits(in); got != want {
			t.Errorf("reverseBits(%d) = %d, want %d", in, got, want)
		}
	}
}
