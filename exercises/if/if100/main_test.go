// if100
// Make the tests pass!

// I AM NOT DONE
//
// direction returns "up", "down" or "idle" for a lift on floor current
// called to floor target. Floors outside 0..maxFloor give "error".
// Practices validating the data and comparing it.
package main_test

import "testing"

func direction(current, target, maxFloor int) string {
	if target > current {
		return "down"
	}
	if target < current {
		return "up"
	}
	return "idle"
}

func TestDirection(t *testing.T) {
	cases := []struct {
		cur, tgt int
		want     string
	}{{0, 5, "up"}, {5, 1, "down"}, {3, 3, "idle"}, {3, 11, "error"}, {-1, 2, "error"}}
	for _, c := range cases {
		if got := direction(c.cur, c.tgt, 10); got != c.want {
			t.Errorf("direction(%d, %d) = %s, want %s", c.cur, c.tgt, got, c.want)
		}
	}
}
