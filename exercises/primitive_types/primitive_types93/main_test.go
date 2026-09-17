// primitive_types93
// Make the tests pass!

// I AM NOT DONE
//
// edge must return the length of the edge of a cube from its volume.
// Practices the root functions of the math package.
package main_test

import (
	"math"
	"testing"
)

func edge(volume float64) float64 {
	return math.Sqrt(volume)
}

func TestEdge(t *testing.T) {
	cases := map[float64]float64{27: 3, 8: 2, 1: 1, 0: 0}
	for in, want := range cases {
		if got := edge(in); got != want {
			t.Errorf("edge(%v) = %v, want %v", in, got, want)
		}
	}
}
