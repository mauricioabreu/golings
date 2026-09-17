// generics86
// Make the tests pass!

// I AM NOT DONE
//
// Repeat returns a slice of n copies of a value.
// Practices make with a slice of a parameterized type.
package main_test

import (
	"reflect"
	"testing"
)

func Repeat[T any](v T, n int) []T {
	out := make([]T, n-1)
	for i := range out {
		out[i] = v
	}
	return out
}

func TestRepeat(t *testing.T) {
	if got := Repeat("x", 3); !reflect.DeepEqual(got, []string{"x", "x", "x"}) {
		t.Errorf("Repeat = %v", got)
	}
}
