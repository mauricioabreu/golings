// range57
// Make the tests pass!

// I AM NOT DONE
//
// pairsWithSum returns every pair of indexes (i < j) whose values add up to target.
// Practices a nested range over a subslice.
package main_test

import (
	"reflect"
	"testing"
)

func pairsWithSum(s []int, target int) [][2]int {
	var out [][2]int
	for i, a := range s {
		for j, b := range s {
			if a+b == target {
				out = append(out, [2]int{i, j})
			}
		}
	}
	return out
}

func TestPairsWithSum(t *testing.T) {
	got := pairsWithSum([]int{1, 5, 3, 3, 2}, 6)
	if !reflect.DeepEqual(got, [][2]int{{0, 1}, {2, 3}}) {
		t.Errorf("pairsWithSum = %v", got)
	}
}
