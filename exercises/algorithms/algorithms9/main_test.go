// algorithms9
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two-dimensional DP. How many paths lead from the top-left corner
// of an r×c grid to the bottom-right one, if you may only move right
// and down?
// Expected asymptotics: O(r·c) time, O(c) space.
package main_test

import "testing"

func uniquePaths(rows, cols int) int {
	return 0
}

func TestUniquePaths(t *testing.T) {
	cases := [][3]int{{3, 7, 28}, {3, 2, 3}, {1, 1, 1}, {1, 10, 1}, {0, 5, 0}, {10, 10, 48620}}
	for _, c := range cases {
		if got := uniquePaths(c[0], c[1]); got != c[2] {
			t.Errorf("uniquePaths(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
