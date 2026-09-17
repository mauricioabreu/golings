// if72
// Make the tests pass!

// I AM NOT DONE
//
// inside must check whether the point (x, y) lies within the rectangle
// from (0, 0) to (w, h), borders included.
// Practices compound conditions over two coordinates.
package main_test

import "testing"

func inside(x, y, w, h int) bool {
	if x >= 0 && x <= w && y >= 0 && y <= w {
		return true
	}
	return false
}

func TestInside(t *testing.T) {
	cases := []struct {
		x, y int
		want bool
	}{{0, 0, true}, {10, 5, true}, {5, 6, false}, {-1, 2, false}, {3, 5, true}}
	for _, c := range cases {
		if got := inside(c.x, c.y, 10, 5); got != c.want {
			t.Errorf("inside(%d, %d) = %v, want %v", c.x, c.y, got, c.want)
		}
	}
}
