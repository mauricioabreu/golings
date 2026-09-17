// if27
// Make the tests pass!

// I AM NOT DONE
//
// close must treat numbers that differ by less than 0.001 as equal.
// Practices comparing floats with a tolerance.
package main_test

import (
	"math"
	"testing"
)

func close(a, b float64) bool {
	if a-b < 0.001 {
		return true
	}
	return false
}

func TestClose(t *testing.T) {
	_ = math.Abs
	if !close(1.0, 1.0005) || !close(1.0005, 1.0) {
		t.Errorf("1.0 and 1.0005 should be close in both orders")
	}
	if close(1.0, 1.1) || close(1.1, 1.0) || close(-5, 5) {
		t.Errorf("distant values should not be close")
	}
}
