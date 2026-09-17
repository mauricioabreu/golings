// range77
// Make the tests pass!

// I AM NOT DONE
//
// splitDiagonal returns the sum of the diagonal elements and the sum of all the others.
// Practices comparing the indexes of two levels of range.
package main_test

import "testing"

func splitDiagonal(m [][]int) (diag, rest int) {
	for i, row := range m {
		for j, v := range row {
			if i == 0 {
				diag += v
			}
			rest += v
		}
	}
	return
}

func TestSplitDiagonal(t *testing.T) {
	d, r := splitDiagonal([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	if d != 15 || r != 30 {
		t.Errorf("splitDiagonal = %d, %d; want 15, 30", d, r)
	}
}
