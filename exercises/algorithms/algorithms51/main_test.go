// algorithms51
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: binary search on the value. An m×n table contains i*j. Find the
// k-th smallest number (counting duplicates, k is 1-based).
// Expected asymptotics: O(m·log(m·n)) time, O(1) space.
package main_test

import "testing"

func findKthNumber(m, n, k int) int {
	return 0
}

func TestFindKthNumber(t *testing.T) {
	cases := [][4]int{{3, 3, 5, 3}, {2, 3, 6, 6}, {1, 1, 1, 1}, {3, 3, 9, 9}, {30000, 30000, 900000000, 900000000}, {45, 12, 471, 312}}
	for _, c := range cases {
		if got := findKthNumber(c[0], c[1], c[2]); got != c[3] {
			t.Errorf("findKthNumber(%d, %d, %d) = %d, want %d", c[0], c[1], c[2], got, c[3])
		}
	}
}
