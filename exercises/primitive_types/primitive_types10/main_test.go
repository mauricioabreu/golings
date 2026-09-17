// primitive_types10
// Make the tests pass!

// I AM NOT DONE
//
// isqrt returns the largest integer r with r*r <= n.
// For large n the result of math.Sqrt may be inexact.
// Practices float64/int conversions and correcting the rounding error.
package main_test

import (
	"math"
	"testing"
)

func isqrt(n int64) int64 {
	r := int64(math.Sqrt(float64(n)))
	return r
}

func TestIsqrt(t *testing.T) {
	cases := map[int64]int64{0: 0, 15: 3, 16: 4, 999999999999999999: 999999999, 1000000000000000000: 1000000000}
	for in, want := range cases {
		if got := isqrt(in); got != want {
			t.Errorf("isqrt(%d) = %d, want %d", in, got, want)
		}
	}
}
