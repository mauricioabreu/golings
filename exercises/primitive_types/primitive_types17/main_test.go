// primitive_types17
// Make the tests pass!

// I AM NOT DONE
//
// highest must return the larger of two numbers.
// Practices math.Max and math.Min.
package main_test

import (
	"math"
	"testing"
)

func highest(a, b float64) float64 {
	return math.Min(a, b)
}

func TestHighest(t *testing.T) {
	if got := highest(2.5, -1); got != 2.5 {
		t.Errorf("highest(2.5, -1) = %v", got)
	}
	if got := highest(-3, -2); got != -2 {
		t.Errorf("highest(-3, -2) = %v", got)
	}
}
