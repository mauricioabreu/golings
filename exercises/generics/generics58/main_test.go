// generics58
// Make the tests pass!

// I AM NOT DONE
//
// Ratio divides two type parameter values as float64.
// The code does not compile: T is constrained by any, and a conversion to float64
// is only allowed for numeric types.
// Conversions depend on the type set of the constraint.
package main_test

import "testing"

type Number interface{ ~int | ~int64 | ~float64 }

func Ratio[T any](a, b T) float64 {
	return float64(a) / float64(b)
}

func TestRatio(t *testing.T) {
	if Ratio(1, 4) != 0.25 || Ratio(3.0, 2.0) != 1.5 {
		t.Errorf("Ratio works incorrectly")
	}
}
