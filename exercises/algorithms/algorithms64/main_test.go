// algorithms64
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: depth-first search with memoization. Find the length of the
// longest path over adjacent cells where the values strictly increase.
// Expected asymptotics: O(r·c) time, O(r·c) space.
package main_test

import "testing"

func longestIncreasingPath(m [][]int) int {
	return 0
}

func TestLongestIncreasingPath(t *testing.T) {
	cases := []struct {
		m    [][]int
		want int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
		{[][]int{{1}}, 1},
		{nil, 0},
		{[][]int{{7, 7}, {7, 7}}, 1},
	}
	for _, c := range cases {
		if got := longestIncreasingPath(c.m); got != c.want {
			t.Errorf("longestIncreasingPath(%v) = %d, want %d", c.m, got, c.want)
		}
	}
}
