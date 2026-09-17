// variables52
// Make the tests pass!

// I AM NOT DONE
//
// midpoint must return the middle of the range [lo, hi] for any int64.
// For large values the result comes out negative.
// Practices overflow when adding large numbers.
package main_test

import (
	"math"
	"testing"
)

func midpoint(lo, hi int64) int64 {
	return (lo + hi) / 2
}

func TestMidpoint(t *testing.T) {
	if got := midpoint(2, 10); got != 6 {
		t.Errorf("midpoint(2, 10) = %d, want 6", got)
	}
	big := int64(math.MaxInt64)
	if got := midpoint(big-10, big); got != big-5 {
		t.Errorf("midpoint(max-10, max) = %d, want %d", got, big-5)
	}
}
