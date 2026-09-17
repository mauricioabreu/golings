// maps79
// Make the tests pass!

// I AM NOT DONE
//
// pairsWithDiff counts the distinct values x for which x+k is in the slice as well (k > 0).
// Practices a set of values and a search for neighbours.
package main_test

import "testing"

func pairsWithDiff(nums []int, k int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}
	count := 0
	for _, v := range nums {
		if set[v-k] {
			count++
		}
	}
	return count
}

func TestPairsWithDiff(t *testing.T) {
	if got := pairsWithDiff([]int{3, 1, 4, 1, 5}, 2); got != 2 {
		t.Errorf("pairsWithDiff = %d, want 2", got)
	}
	if got := pairsWithDiff([]int{1, 3, 1, 3}, 2); got != 1 {
		t.Errorf("pairsWithDiff = %d, want 1", got)
	}
}
