// primitive_types11
// Make the tests pass!

// I AM NOT DONE
//
// roundCents rounds an amount to whole cents the banker's way:
// a half goes to the even value (0.5 -> 0, 1.5 -> 2, 2.5 -> 2).
// math.Round sends a half away from zero; math.RoundToEven exists too.
package main_test

import (
	"math"
	"testing"
)

func roundCents(x float64) float64 {
	return math.Round(x)
}

func TestRoundCents(t *testing.T) {
	cases := map[float64]float64{0.5: 0, 1.5: 2, 2.5: 2, 3.5: 4, -2.5: -2, 2.6: 3}
	for in, want := range cases {
		if got := roundCents(in); got != want {
			t.Errorf("roundCents(%v) = %v, want %v", in, got, want)
		}
	}
}
