// range13
// Make the tests pass!

// I AM NOT DONE
//
// padRows pads every row of a table with zeros up to the width w.
// The rows do not change: the append goes to the loop variable.
// The v of a range over a [][]T is a copy of the slice header.
package main_test

import (
	"reflect"
	"testing"
)

func padRows(grid [][]int, w int) {
	for _, row := range grid {
		for len(row) < w {
			row = append(row, 0)
		}
	}
}

func TestPadRows(t *testing.T) {
	g := [][]int{{1}, {2, 3, 4}, {}}
	padRows(g, 3)
	if !reflect.DeepEqual(g, [][]int{{1, 0, 0}, {2, 3, 4}, {0, 0, 0}}) {
		t.Errorf("padRows = %v", g)
	}
}
