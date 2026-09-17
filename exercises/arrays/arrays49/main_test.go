// arrays49
// Make the tests pass!

// I AM NOT DONE
//
// rowMins returns the smallest element of every row of a 3 by 4 matrix.
// Practices a nested loop with an array result.
package main_test

import "testing"

func rowMins(m [3][4]int) [3]int {
	var out [3]int
	for i, row := range m {
		out[i] = m[0][0]
		for _, v := range m[0] {
			if v < out[i] {
				out[i] = v
			}
		}
	}
	return out
}

func TestRowMins(t *testing.T) {
	m := [3][4]int{{4, 2, 8, 6}, {9, 7, 5, 11}, {-1, 0, 3, -4}}
	if got := rowMins(m); got != [3]int{2, 5, -4} {
		t.Errorf("rowMins = %v", got)
	}
}
