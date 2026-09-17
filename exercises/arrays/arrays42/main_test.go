// arrays42
// Make the tests pass!

// I AM NOT DONE
//
// diagonals returns the sums of the main and the anti diagonal of a 4 by 4 matrix.
// Practices the indexes [i][n-1-i].
package main_test

import "testing"

func diagonals(m [4][4]int) (int, int) {
	main, anti := 0, 0
	n := len(m)
	for i := 0; i < n; i++ {
		main += m[i][n-1-i]
		anti += m[n-1-i][n-1-i]
	}
	return main, anti
}

func TestDiagonals(t *testing.T) {
	m := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	d, a := diagonals(m)
	if d != 34 || a != 34 {
		t.Errorf("diagonals = %d, %d; want 34, 34", d, a)
	}
	m[0][0] = 0
	if d, _ := diagonals(m); d != 33 {
		t.Errorf("main diagonal = %d, want 33", d)
	}
}
