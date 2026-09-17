// if56
// Make the tests pass!

// I AM NOT DONE
//
// inRange must check that x lies in the closed range [lo, hi].
// Practices compound conditions with bounds.
package main_test

import "testing"

func inRange(x, lo, hi int) bool {
	if x > lo && x < hi {
		return true
	}
	return false
}

func TestInRange(t *testing.T) {
	cases := []struct {
		x    int
		want bool
	}{{1, true}, {5, true}, {10, true}, {0, false}, {11, false}}
	for _, c := range cases {
		if got := inRange(c.x, 1, 10); got != c.want {
			t.Errorf("inRange(%d, 1, 10) = %v, want %v", c.x, got, c.want)
		}
	}
}
