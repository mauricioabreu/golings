// arrays11
// Make the tests pass!

// I AM NOT DONE
//
// validRow checks that the digits 1-9 do not repeat in a sudoku row, 0 being an empty cell.
// Practices a [10]bool counter array.
package main_test

import "testing"

func validRow(row [9]int) bool {
	var seen [10]bool
	for _, v := range row {
		if seen[v] {
			return false
		}
	}
	return true
}

func TestValidRow(t *testing.T) {
	if !validRow([9]int{5, 3, 0, 0, 7, 0, 0, 0, 0}) {
		t.Errorf("row with blanks should be valid")
	}
	if validRow([9]int{5, 3, 5, 0, 7, 0, 0, 0, 0}) {
		t.Errorf("row with duplicate 5 should be invalid")
	}
}
