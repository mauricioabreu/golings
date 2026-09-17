// variables75
// Make the tests pass!

// I AM NOT DONE
//
// almostEqual must treat 0.1+0.2 and 0.3 as equal.
// An exact comparison of floating point numbers does not work here.
// Practices float64 rounding error and comparing against an epsilon.
package main_test

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func almostEqual(a, b float64) bool {
	_ = math.Abs
	return a == b
}

func TestAlmostEqual(t *testing.T) {
	x, y := 0.1, 0.2
	if !almostEqual(x+y, 0.3) {
		t.Errorf("0.1+0.2 should be almost equal to 0.3")
	}
	if almostEqual(0.3, 0.31) {
		t.Errorf("0.3 and 0.31 should not be almost equal")
	}
}
