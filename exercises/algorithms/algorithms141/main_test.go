// algorithms141
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: bit manipulation. Count in how many bits two non-negative
// numbers differ.
// Expected asymptotics: O(1) time, O(1) space.
package main_test

import "testing"

func hammingDistance(a, b int) int {
	return 0
}

func TestHammingDistance(t *testing.T) {
	cases := [][3]int{{1, 4, 2}, {3, 1, 1}, {0, 0, 0}, {0, 1 << 40, 1}, {255, 0, 8}}
	for _, c := range cases {
		if got := hammingDistance(c[0], c[1]); got != c[2] {
			t.Errorf("hammingDistance(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
