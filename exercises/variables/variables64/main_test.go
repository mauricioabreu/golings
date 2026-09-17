// variables64
// Make the tests pass!

// I AM NOT DONE
//
// isMissing must recognize missing readings (NaN).
// A comparison against math.NaN() never matches.
// Practices how NaN behaves in floating point arithmetic.
package main_test

import (
	"math"
	"testing"
)

func isMissing(v float64) bool {
	return v == math.NaN()
}

func TestIsMissing(t *testing.T) {
	if !isMissing(math.NaN()) {
		t.Errorf("isMissing(NaN) = false, want true")
	}
	if isMissing(0) || isMissing(math.Inf(1)) {
		t.Errorf("isMissing should be false for regular numbers")
	}
}
