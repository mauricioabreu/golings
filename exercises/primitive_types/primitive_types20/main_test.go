// primitive_types20
// Make the tests pass!

// I AM NOT DONE
//
// compound computes the value of a deposit: amount * (1 + rate)^years.
// Practices math.Pow.
package main_test

import (
	"math"
	"testing"
)

func compound(amount, rate float64, years int) float64 {
	return amount * math.Pow(float64(years), 1+rate)
}

func TestCompound(t *testing.T) {
	if got := compound(1000, 0.1, 2); math.Abs(got-1210) > 1e-9 {
		t.Errorf("compound(1000, 0.1, 2) = %v, want 1210", got)
	}
	if got := compound(500, 0.5, 0); got != 500 {
		t.Errorf("compound(500, 0.5, 0) = %v, want 500", got)
	}
}
