// algorithms76
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: bit tricks. Count the set bits in an unsigned number without
// using math/bits.
// Expected asymptotics: O(k) time (k is the number of set bits), O(1) space.
package main_test

import "testing"

func hammingWeight(x uint64) int {
	return 0
}

func TestHammingWeight(t *testing.T) {
	cases := map[uint64]int{0: 0, 1: 1, 11: 3, 128: 1, 1 << 63: 1, ^uint64(0): 64}
	for in, want := range cases {
		if got := hammingWeight(in); got != want {
			t.Errorf("hammingWeight(%d) = %d, want %d", in, got, want)
		}
	}
}
