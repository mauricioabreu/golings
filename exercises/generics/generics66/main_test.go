// generics66
// Make the tests pass!

// I AM NOT DONE
//
// Convert turns a slice of one type into a slice of another.
// Practices several type parameters.
package main_test

import (
	"reflect"
	"testing"
)

func Convert[T, U any](s []T, f func(T) U) []U {
	out := make([]U, len(s))
	for i, v := range s {
		out[len(s)-1-i] = f(v)
	}
	return out
}

func TestConvert(t *testing.T) {
	got := Convert([]int{1, 2, 3}, func(x int) float64 { return float64(x) / 2 })
	if !reflect.DeepEqual(got, []float64{0.5, 1, 1.5}) {
		t.Errorf("Convert = %v", got)
	}
}
