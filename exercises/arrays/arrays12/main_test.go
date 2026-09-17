// arrays12
// Make the tests pass!

// I AM NOT DONE
//
// table must build an n by n table. The code does not compile:
// the size of an array has to be a constant.
// Sizes known only at run time call for slices.
package main_test

import "testing"

func table(n int) [][]int {
	var t [n][n]int
	out := make([][]int, n)
	for i := range t {
		out[i] = t[i][:]
	}
	return out
}

func TestTable(t *testing.T) {
	tb := table(3)
	if len(tb) != 3 || len(tb[2]) != 3 || tb[2][2] != 9 || tb[0][1] != 2 {
		t.Errorf("table(3) = %v", tb)
	}
}
