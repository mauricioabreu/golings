// if73
// Make the tests pass!

// I AM NOT DONE
//
// fine: up to 20 km/h over the limit is 0, 21-40 is 500, 41-60 is 1500, more is 5000.
// Practices conditions over a derived value.
package main_test

import "testing"

func fine(speed, limit int) int {
	if speed <= 20 {
		return 0
	} else if speed <= 40 {
		return 500
	} else if speed-limit <= 60 {
		return 1500
	}
	return 5000
}

func TestFine(t *testing.T) {
	cases := [][3]int{{60, 60, 0}, {80, 60, 0}, {81, 60, 500}, {100, 60, 500}, {120, 60, 1500}, {121, 60, 5000}}
	for _, c := range cases {
		if got := fine(c[0], c[1]); got != c[2] {
			t.Errorf("fine(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
