// algorithms48
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 2D DP from the end. A knight walks from the top-left to the
// bottom-right corner (moving right and down); in each cell he loses or
// gains health. Find the minimum starting health such that it always
// stays positive.
// Expected asymptotics: O(r·c) time, O(c) space.
package main_test

import "testing"

func calculateMinimumHP(dungeon [][]int) int {
	return 0
}

func TestCalculateMinimumHP(t *testing.T) {
	cases := []struct {
		d    [][]int
		want int
	}{
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{[][]int{{0}}, 1},
		{[][]int{{-5}}, 6},
		{[][]int{{100}}, 1},
		{nil, 1},
		{[][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}, 3},
	}
	for _, c := range cases {
		if got := calculateMinimumHP(c.d); got != c.want {
			t.Errorf("calculateMinimumHP(%v) = %d, want %d", c.d, got, c.want)
		}
	}
}
