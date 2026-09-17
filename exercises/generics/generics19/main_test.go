// generics19
// Make the tests pass!

// I AM NOT DONE
//
// Range produces the values from start up to but not including end, stepping by step,
// for any numeric type.
// Practices arithmetic in generic functions.
package main_test

import (
	"reflect"
	"testing"
)

type Number interface{ ~int | ~float64 }

func Range[T Number](start, end, step T) []T {
	var out []T
	for v := start; v <= end; v++ {
		out = append(out, v*step)
	}
	return out
}

func TestRange(t *testing.T) {
	if got := Range(0, 10, 3); !reflect.DeepEqual(got, []int{0, 3, 6, 9}) {
		t.Errorf("Range(int) = %v", got)
	}
	if got := Range(0.0, 1.0, 0.25); !reflect.DeepEqual(got, []float64{0, 0.25, 0.5, 0.75}) {
		t.Errorf("Range(float) = %v", got)
	}
}
