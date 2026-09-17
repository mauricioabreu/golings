// slices36
// Make the tests pass!

// I AM NOT DONE
//
// rotate turns a rectangular matrix 90 degrees clockwise.
// Practices the dimensions of the result for a non-square matrix.
package main_test

import (
	"reflect"
	"testing"
)

func rotate(m [][]int) [][]int {
	rows, cols := len(m), len(m[0])
	out := make([][]int, rows)
	for i := range out {
		out[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			out[i][cols-1-j] = m[i][j]
		}
	}
	return out
}

func TestRotate(t *testing.T) {
	got := rotate([][]int{{1, 2, 3}, {4, 5, 6}})
	if !reflect.DeepEqual(got, [][]int{{4, 1}, {5, 2}, {6, 3}}) {
		t.Errorf("rotate = %v", got)
	}
}
