// algorithms121
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: backtracking with constraint checking. Fill a 9×9 board (empty
// cells are '.') with digits 1–9 so that no digit repeats in any row,
// column, or 3×3 square. A solution exists and is unique.
// Expected asymptotics: O(9^k) time, O(1) space.
package main_test

import "testing"

func solveSudoku(board *[9][9]byte) bool {
	return false
}

func TestSolveSudoku(t *testing.T) {
	rows := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	var board [9][9]byte
	for r, row := range rows {
		copy(board[r][:], row)
	}
	if !solveSudoku(&board) {
		t.Fatalf("sudoku was not solved")
	}
	for r := 0; r < 9; r++ {
		var seen [10]bool
		for c := 0; c < 9; c++ {
			v := board[r][c]
			if v < '1' || v > '9' || seen[v-'0'] {
				t.Fatalf("row %d is invalid: %q", r, board[r])
			}
			seen[v-'0'] = true
		}
	}
	if board[0][2] != '4' || board[8][8] != '9' {
		t.Errorf("unexpected solution: %q", board[0])
	}
}
