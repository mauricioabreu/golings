// generics35
// Make the tests pass!

// I AM NOT DONE
//
// average has to compute the mean in float64, calling Sum with an explicit type argument.
// Practices explicit type arguments.
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Sum[T Number](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

func average() float64 {
	return float64(Sum[int](1, 2) / 2)
}

func TestAverage(t *testing.T) {
	if average() != 1.5 {
		t.Errorf("average = %v", average())
	}
}
