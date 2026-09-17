// switch27
// Make the tests pass!

// I AM NOT DONE
//
// pieceValue returns the value of a chess piece: a pawn 1, a knight and a bishop 3,
// a rook 5, a queen 9.
// Practices grouping values in one branch.
package main_test

import "testing"

func pieceValue(p byte) int {
	switch p {
	case 'P':
		return 1
	case 'N':
		return 3
	case 'R':
		return 5
	case 'Q':
		return 9
	}
	return 0
}

func TestPieceValue(t *testing.T) {
	cases := map[byte]int{'P': 1, 'N': 3, 'B': 3, 'R': 5, 'Q': 9, 'K': 0}
	for in, want := range cases {
		if got := pieceValue(in); got != want {
			t.Errorf("pieceValue(%c) = %d, want %d", in, got, want)
		}
	}
}
