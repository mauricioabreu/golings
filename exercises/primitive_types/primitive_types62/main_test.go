// primitive_types62
// Make the tests pass!

// I AM NOT DONE
//
// mean returns the average value, and must return 0 for an empty slice.
// In floating point 0/0 does not panic, it gives NaN.
// Division by zero behaves differently for integers and for float64.
package main_test

import (
	"math"
	"testing"
)

func mean(xs []float64) float64 {
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func TestMean(t *testing.T) {
	if got := mean([]float64{1, 2, 3}); got != 2 {
		t.Errorf("mean(1,2,3) = %v", got)
	}
	if got := mean(nil); got != 0 || math.IsNaN(got) {
		t.Errorf("mean(nil) = %v, want 0", got)
	}
}
