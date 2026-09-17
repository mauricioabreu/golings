// if97
// Make the tests pass!

// I AM NOT DONE
//
// quadrant returns the quadrant number (1-4) of a point, or 0 when it sits on an axis.
// Practices conditions over two variables.
package main_test

import "testing"

func quadrant(x, y int) int {
	if x > 0 && y > 0 {
		return 1
	} else if x < 0 && y > 0 {
		return 3
	} else if x < 0 && y < 0 {
		return 2
	}
	return 4
}

func TestQuadrant(t *testing.T) {
	cases := []struct{ x, y, want int }{
		{1, 1, 1}, {-1, 1, 2}, {-1, -1, 3}, {1, -1, 4}, {0, 5, 0}, {3, 0, 0},
	}
	for _, c := range cases {
		if got := quadrant(c.x, c.y); got != c.want {
			t.Errorf("quadrant(%d, %d) = %d, want %d", c.x, c.y, got, c.want)
		}
	}
}
