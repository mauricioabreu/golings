// primitive_types7
// Make the tests pass!

// I AM NOT DONE
//
// roundTo rounds a number to n decimal places.
// Practices math.Pow10 and math.Round.
package main_test

import (
	"math"
	"testing"
)

func roundTo(x float64, n int) float64 {
	p := math.Pow10(n - 1)
	return math.Floor(x*p) / p
}

func TestRoundTo(t *testing.T) {
	cases := []struct {
		x    float64
		n    int
		want float64
	}{{3.14159, 2, 3.14}, {2.675, 1, 2.7}, {-1.25, 1, -1.3}, {1234.5, 0, 1235}}
	for _, c := range cases {
		if got := roundTo(c.x, c.n); got != c.want {
			t.Errorf("roundTo(%v, %d) = %v, want %v", c.x, c.n, got, c.want)
		}
	}
}
