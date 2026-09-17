// range98
// Make the tests pass!

// I AM NOT DONE
//
// validRows returns the indexes of the rows holding no negative numbers.
// Practices continue with a label on the outer loop.
package main_test

import (
	"reflect"
	"testing"
)

func validRows(g [][]int) []int {
	var out []int
	for i, row := range g {
		for _, v := range row {
			if v < 0 {
				continue
			}
		}
		out = append(out, i)
	}
	return out
}

func TestValidRows(t *testing.T) {
	got := validRows([][]int{{1, 2}, {3, -1}, {}, {-5}})
	if !reflect.DeepEqual(got, []int{0, 2}) {
		t.Errorf("validRows = %v", got)
	}
}
