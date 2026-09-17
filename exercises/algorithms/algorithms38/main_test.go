// algorithms38
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: breadth-first traversal by levels. Cells: 0 — empty, 1 — a fresh
// orange, 2 — a rotten one. Each minute the rot spreads to neighbors.
// Return the minutes until all rot, or -1 if that is impossible.
// Expected asymptotics: O(r·c) time, O(r·c) space.
package main_test

import "testing"

func orangesRotting(grid [][]int) int {
	return 0
}

func TestOrangesRotting(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{0}}, 0},
		{[][]int{{1}}, -1},
	}
	for _, c := range cases {
		if got := orangesRotting(c.grid); got != c.want {
			t.Errorf("orangesRotting = %d, want %d", got, c.want)
		}
	}
}
