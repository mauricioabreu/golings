// arrays36
// Make the tests pass!

// I AM NOT DONE
//
// sumColumns adds up the columns of a 2 by 3 matrix (2 rows, 3 columns).
// The row and column indexes are muddled up.
// len(m) and len(m[0]) on a non-square array.
package main_test

import "testing"

func sumColumns(m [2][3]int) [3]int {
	var out [3]int
	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m); j++ {
			out[j] += m[i][j]
		}
	}
	return out
}

func TestSumColumns(t *testing.T) {
	_ = sumColumns
	m := [2][3]int{{1, 2, 3}, {10, 20, 30}}
	if got := sumColumns(m); got != [3]int{11, 22, 33} {
		t.Errorf("sumColumns = %v", got)
	}
}
