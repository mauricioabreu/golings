// arrays45
// Make the tests pass!

// I AM NOT DONE
//
// samePoint compares [2]float64 points and must treat points with
// missing (NaN) coordinates in the same places as equal.
// == on arrays compares the elements, and NaN != NaN.
package main_test

import (
	"math"
	"testing"
)

func samePoint(a, b [2]float64) bool {
	return a == b
}

func TestSamePoint(t *testing.T) {
	nan := math.NaN()
	if !samePoint([2]float64{1, nan}, [2]float64{1, nan}) {
		t.Errorf("points with NaN in the same place should match")
	}
	if samePoint([2]float64{1, nan}, [2]float64{nan, 1}) || samePoint([2]float64{1, 2}, [2]float64{1, 3}) {
		t.Errorf("different points should not match")
	}
}
