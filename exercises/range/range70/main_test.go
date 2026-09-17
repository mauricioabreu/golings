// range70
// Make the tests pass!

// I AM NOT DONE
//
// gridSum adds up every element of a two dimensional slice.
// Practices a nested range.
package main_test

import "testing"

func gridSum(g [][]int) int {
	sum := 0
	for _, row := range g {
		for _, v := range g[0] {
			sum += v
		}
	}
	return sum
}

func TestGridSum(t *testing.T) {
	if got := gridSum([][]int{{1, 2}, {3, 4, 5}}); got != 15 {
		t.Errorf("gridSum = %d, want 15", got)
	}
}
