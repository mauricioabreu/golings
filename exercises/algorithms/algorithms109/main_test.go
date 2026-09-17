// algorithms109
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 2D dynamic programming. Find the minimum sum of numbers on a
// path from the top-left to the bottom-right corner of a grid, moving
// right and down.
// Expected asymptotics: O(r·c) time, O(c) space.
package main_test

import "testing"

func minPathSum(grid [][]int) int {
	return 0
}

func TestMinPathSum(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{[][]int{{5}}, 5},
		{nil, 0},
		{[][]int{{1, 2}, {1, 1}}, 3},
	}
	for _, c := range cases {
		if got := minPathSum(c.grid); got != c.want {
			t.Errorf("minPathSum(%v) = %d, want %d", c.grid, got, c.want)
		}
	}
}
