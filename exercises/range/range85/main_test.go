// range85
// Make the tests pass!

// I AM NOT DONE
//
// columnMax returns the maximum of every column of a rectangular table.
// Practices a range over rows and columns with a shared result.
package main_test

import (
	"reflect"
	"testing"
)

func columnMax(g [][]int) []int {
	if len(g) == 0 {
		return nil
	}
	out := append([]int(nil), g[0]...)
	for _, row := range g[1:] {
		for c := range row {
			if row[c] > out[0] {
				out[0] = row[c]
			}
		}
	}
	return out
}

func TestColumnMax(t *testing.T) {
	got := columnMax([][]int{{1, -5, 3}, {4, -6, 0}, {2, -1, 9}})
	if !reflect.DeepEqual(got, []int{4, -1, 9}) {
		t.Errorf("columnMax = %v", got)
	}
}
