// primitive_types95
// Make the tests pass!

// I AM NOT DONE
//
// lowestStart must return negative infinity, the starting value
// for a search for a maximum.
// Practices math.Inf and the sign of an infinity.
package main_test

import (
	"math"
	"testing"
)

func lowestStart() float64 {
	return math.Inf(1)
}

func TestLowestStart(t *testing.T) {
	if got := lowestStart(); !math.IsInf(got, -1) {
		t.Errorf("lowestStart() = %v, want -Inf", got)
	}
}
