// generics38
// Make the tests pass!

// I AM NOT DONE
//
// The code tries to declare a slice of values of the Number constraint.
// It does not compile: an interface with a type set may only be used as a constraint.
// Practices the difference between a constraint and an ordinary interface.
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Largest(vals []Number) Number {
	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

func TestLargest(t *testing.T) {
	if Largest([]int{3, 8, 1}) != 8 || Largest([]float64{0.5, 0.25}) != 0.5 {
		t.Errorf("Largest works incorrectly")
	}
}
