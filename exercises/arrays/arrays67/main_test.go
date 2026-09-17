// arrays67
// Make the tests pass!

// I AM NOT DONE
//
// board builds an 8 by 8 board where true is a black square; the square a1 ([7][0]) is black.
// Practices filling a two dimensional array from a formula.
package main_test

import "testing"

func board() [8][8]bool {
	var b [8][8]bool
	for r := range b {
		b[r][0] = r%2 == 0
	}
	return b
}

func TestBoard(t *testing.T) {
	b := board()
	if !b[7][0] || b[7][1] || b[0][0] || !b[0][1] || !b[3][4] {
		t.Errorf("board colors are wrong")
	}
}
