// primitive_types73
// Make the tests pass!

// I AM NOT DONE
//
// nearlyEqual compares numbers with a relative precision of 1e-9, so that it works
// for very large and very small values alike.
// Practices comparing floats with their scale in mind.
package main_test

import (
	"math"
	"testing"
)

func nearlyEqual(a, b float64) bool {
	if a == b {
		return true
	}
	return math.Abs(a-b) < 1e-9
}

func TestNearlyEqual(t *testing.T) {
	if !nearlyEqual(1e20, 1e20+1e5) {
		t.Errorf("1e20 and 1e20+1e5 should be nearly equal")
	}
	if nearlyEqual(1e-12, 2e-12) {
		t.Errorf("1e-12 and 2e-12 differ by a factor of 2")
	}
	if !nearlyEqual(0.1+0.2, 0.3) {
		t.Errorf("0.1+0.2 should be nearly 0.3")
	}
}
