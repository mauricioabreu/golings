// slices8
// Make the tests pass!

// I AM NOT DONE
//
// newGrid builds a rows by cols grid filled with the value v.
// Practices make for the outer slice and for every row.
package main_test

import "testing"

func newGrid(rows, cols, v int) [][]int {
	g := make([][]int, rows)
	for r := range g {
		for c := 0; c < cols; c++ {
			g[r] = append(g[r], c)
		}
	}
	return g
}

func TestNewGrid(t *testing.T) {
	g := newGrid(2, 3, 7)
	if len(g) != 2 || len(g[1]) != 3 || g[1][2] != 7 || g[0][0] != 7 {
		t.Errorf("newGrid = %v", g)
	}
}
