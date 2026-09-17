// structs72
// Make the tests pass!

// I AM NOT DONE
//
// Coordinate structs are used as map keys.
// Practices comparable structs as keys.
package main_test

import "testing"

type Cell struct{ Row, Col int }

func mark(board map[Cell]string, r, c int, v string) {
	board[Cell{c, r}] = v
}

func TestMark(t *testing.T) {
	b := map[Cell]string{}
	mark(b, 0, 2, "X")
	if b[Cell{0, 2}] != "X" {
		t.Errorf("board = %v", b)
	}
}
