// generics92
// Make the tests pass!

// I AM NOT DONE
//
// AddMatrix adds two matrices of any numeric type.
// Practices generic two dimensional slices.
package main_test

import (
	"reflect"
	"testing"
)

type Number interface{ ~int | ~float64 }

func AddMatrix[T Number](a, b [][]T) [][]T {
	out := make([][]T, len(a))
	for i := range a {
		out[i] = a[i]
		for j := range a[i] {
			out[i][j] += b[i][j]
		}
	}
	return out
}

func TestAddMatrix(t *testing.T) {
	a := [][]float64{{1, 2}}
	got := AddMatrix(a, [][]float64{{0.5, 0.5}})
	if !reflect.DeepEqual(got, [][]float64{{1.5, 2.5}}) {
		t.Errorf("AddMatrix = %v", got)
	}
	if a[0][0] != 1 {
		t.Errorf("input modified: %v", a)
	}
}
