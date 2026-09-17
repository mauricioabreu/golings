// primitive_types88
// Make the tests pass!

// I AM NOT DONE
//
// frac360 must return the remainder of an angle divided by 360, the angle being non-negative.
// The code does not compile: the % operator is not defined for float64.
// Practices math.Mod.
package main_test

import (
	"math"
	"testing"
)

func frac360(angle float64) float64 {
	return angle % 360
}

func TestFrac360(t *testing.T) {
	_ = math.Mod
	cases := map[float64]float64{370.5: 10.5, 720: 0, 45: 45}
	for in, want := range cases {
		if got := frac360(in); got != want {
			t.Errorf("frac360(%v) = %v, want %v", in, got, want)
		}
	}
}
