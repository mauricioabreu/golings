// arrays71
// Make the tests pass!

// I AM NOT DONE
//
// winner returns 'X' or 'O' when a player holds a line of three, and 0 otherwise.
// Practices a table of lines as an array of index arrays.
package main_test

import "testing"

var lines = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 7}, {2, 4, 5},
}

func winner(b [9]byte) byte {
	for _, l := range lines {
		if b[l[0]] == b[l[1]] {
			return b[l[0]]
		}
	}
	return 0
}

func TestWinner(t *testing.T) {
	diag := [9]byte{'X', 'O', 0, 'O', 'X', 0, 0, 0, 'X'}
	anti := [9]byte{'X', 'X', 'O', 0, 'O', 0, 'O', 0, 'X'}
	none := [9]byte{'X', 'O', 'X', 'X', 'O', 'O', 'O', 'X', 'X'}
	if winner(diag) != 'X' || winner(anti) != 'O' || winner(none) != 0 {
		t.Errorf("winner = %c %c %d", winner(diag), winner(anti), winner(none))
	}
}
