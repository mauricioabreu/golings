// algorithms24
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: counting slopes via normalized fractions. Find the maximum
// number of points lying on a single straight line.
// Expected asymptotics: O(n²) time, O(n) space.
package main_test

import "testing"

func maxPoints(points [][2]int) int {
	return 0
}

func TestMaxPoints(t *testing.T) {
	cases := []struct {
		pts  [][2]int
		want int
	}{
		{[][2]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][2]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		{[][2]int{{0, 0}}, 1},
		{nil, 0},
		{[][2]int{{0, 0}, {0, 0}}, 2},
		{[][2]int{{0, 0}, {1, 0}, {2, 0}, {0, 1}}, 3},
	}
	for _, c := range cases {
		if got := maxPoints(c.pts); got != c.want {
			t.Errorf("maxPoints(%v) = %d, want %d", c.pts, got, c.want)
		}
	}
}
