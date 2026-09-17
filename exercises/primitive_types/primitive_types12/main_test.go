// primitive_types12
// Make the tests pass!

// I AM NOT DONE
//
// distance must return the distance between (x1, y1) and (x2, y2).
// Practices math.Hypot.
package main_test

import (
	"math"
	"testing"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x2-x1, y2-x1)
}

func TestDistance(t *testing.T) {
	if got := distance(1, 1, 4, 5); got != 5 {
		t.Errorf("distance = %v, want 5", got)
	}
	if got := distance(2, 0, 2, 3); got != 3 {
		t.Errorf("distance = %v, want 3", got)
	}
}
