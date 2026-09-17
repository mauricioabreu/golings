// variables65
// Make the tests pass!

// I AM NOT DONE
//
// hypot must return the length of the hypotenuse for integer sides.
// The code does not compile: math.Sqrt takes a float64.
// Practices explicit numeric conversions.
package main_test

import (
	"math"
	"testing"
)

func hypot(a, b int) float64 {
	sq := a*a + b*b
	return math.Sqrt(sq)
}

func TestHypot(t *testing.T) {
	if got := hypot(3, 4); got != 5 {
		t.Errorf("hypot(3, 4) = %v, want 5", got)
	}
}
