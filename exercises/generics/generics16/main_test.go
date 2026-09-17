// generics16
// Make the tests pass!

// I AM NOT DONE
//
// Half halves a value by multiplying it by 0.5. The code does not compile:
// the type set holds integer types, which cannot represent 0.5.
// Constants are checked against every type of a constraint.
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Half[T Number](v T) T {
	return v * 0.5
}

func TestHalf(t *testing.T) {
	if Half(10) != 5 || Half(3.0) != 1.5 {
		t.Errorf("Half works incorrectly")
	}
}
