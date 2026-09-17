// generics13
// Make the tests pass!

// I AM NOT DONE
//
// Total is called on a slice of int32, and the constraint does not list that type.
// Practices a union of types in a constraint.
package main_test

import "testing"

type Integer interface {
	~int | ~int64
}

func Total[T Integer](s []T) T {
	var t T
	for _, v := range s {
		t += v
	}
	return t
}

func TestTotal(t *testing.T) {
	if Total([]int32{1, 2, 3}) != 6 {
		t.Errorf("Total = %d", Total([]int32{1, 2, 3}))
	}
}
