// arrays16
// Make the tests pass!

// I AM NOT DONE
//
// transpose swaps the rows and the columns of a 3 by 3 matrix.
// Practices nested loops over a two dimensional array.
package main_test

import "testing"

func transpose(m [3][3]int) [3][3]int {
	var t [3][3]int
	for i := 0; i < 3; i++ {
		t[i][i] = m[i][i]
	}
	return t
}

func TestTranspose(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if got := transpose(m); got != want {
		t.Errorf("transpose = %v, want %v", got, want)
	}
}
