// slices103
// Make the tests pass!

// I AM NOT DONE
//
// transpose swaps the rows and the columns of a rectangular matrix.
// Practices building slices of the right sizes.
package main_test

import (
	"reflect"
	"testing"
)

func transpose(m [][]int) [][]int {
	if len(m) == 0 {
		return nil
	}
	out := make([][]int, len(m))
	for i := range m {
		out[i] = make([]int, len(m[0]))
		for j := range m[i] {
			out[i][j] = m[i][j]
		}
	}
	return out
}

func TestTranspose(t *testing.T) {
	got := transpose([][]int{{1, 2, 3}, {4, 5, 6}})
	if !reflect.DeepEqual(got, [][]int{{1, 4}, {2, 5}, {3, 6}}) {
		t.Errorf("transpose = %v", got)
	}
}
