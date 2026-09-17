// arrays57
// Make the tests pass!

// I AM NOT DONE
//
// cell must return the value in row r and column c.
// Practices indexing a two dimensional array.
package main_test

import "testing"

func cell(grid [2][3]int, r, c int) int {
	return grid[c][r]
}

func TestCell(t *testing.T) {
	g := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	if got := cell(g, 1, 2); got != 6 {
		t.Errorf("cell(1, 2) = %d, want 6", got)
	}
	if got := cell(g, 0, 1); got != 2 {
		t.Errorf("cell(0, 1) = %d, want 2", got)
	}
}
