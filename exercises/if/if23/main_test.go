// if23
// Make the tests pass!

// I AM NOT DONE
//
// overlap must return true when the ranges [a1, a2] and [b1, b2] overlap.
// Practices the condition for two ranges to meet.
package main_test

import "testing"

func overlap(a1, a2, b1, b2 int) bool {
	if a1 <= b2 || b1 <= a2 {
		return true
	}
	return false
}

func TestOverlap(t *testing.T) {
	cases := []struct {
		a1, a2, b1, b2 int
		want           bool
	}{{1, 5, 4, 8, true}, {1, 3, 4, 6, false}, {5, 9, 1, 2, false}, {1, 10, 3, 4, true}, {1, 2, 2, 3, true}}
	for _, c := range cases {
		if got := overlap(c.a1, c.a2, c.b1, c.b2); got != c.want {
			t.Errorf("overlap(%v) = %v, want %v", c, got, c.want)
		}
	}
}
