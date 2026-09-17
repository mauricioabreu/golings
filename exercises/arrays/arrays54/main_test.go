// arrays54
// Make the tests pass!

// I AM NOT DONE
//
// Board.Row returns a row of the board, and the caller wants to change it.
// The changes are lost because a copy of the array is returned.
// Returning an array from a function copies it.
package main_test

import "testing"

type Board struct {
	cells [2][4]int
}

func (b *Board) Row(i int) [4]int {
	return b.cells[i]
}

func TestRow(t *testing.T) {
	var b Board
	row := b.Row(1)
	row[3] = 42
	if b.cells[1][3] != 42 {
		t.Errorf("cells[1][3] = %d, want 42", b.cells[1][3])
	}
}
