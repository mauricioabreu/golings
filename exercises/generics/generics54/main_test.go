// generics54
// Make the tests pass!

// I AM NOT DONE
//
// Total adds up values; called on a []int8 the sum overflows.
// Arithmetic in generic code happens in the type of the argument.
package main_test

import "testing"

type Integer interface{ ~int8 | ~int16 | ~int | ~int64 }

func Total[T Integer](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func TestTotal(t *testing.T) {
	if got := Total([]int8{100, 100, 100}); got != 300 {
		t.Errorf("Total = %d, want 300", got)
	}
}
