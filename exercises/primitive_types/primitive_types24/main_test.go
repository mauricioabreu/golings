// primitive_types24
// Make the tests pass!

// I AM NOT DONE
//
// normalize maps an angle into the range [0, 360).
// For negative angles the result is negative.
// math.Mod keeps the sign of the dividend.
package main_test

import (
	"math"
	"testing"
)

func normalize(deg float64) float64 {
	r := math.Mod(deg, 360)
	return r
}

func TestNormalize(t *testing.T) {
	cases := map[float64]float64{30: 30, 370: 10, -30: 330, -720: 0, 359.5: 359.5}
	for in, want := range cases {
		if got := normalize(in); got != want {
			t.Errorf("normalize(%v) = %v, want %v", in, got, want)
		}
	}
}
