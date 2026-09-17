// functions100
// Make the tests pass!

// I AM NOT DONE
//
// lowerBound must return the index of the first element >= target in a sorted slice.
// Practices sort.Search and monotone predicates.
package main_test

import (
	"sort"
	"testing"
)

func lowerBound(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
}

func TestLowerBound(t *testing.T) {
	nums := []int{1, 3, 3, 5, 8}
	cases := map[int]int{3: 1, 4: 3, 0: 0, 9: 5, 8: 4}
	for target, want := range cases {
		if got := lowerBound(nums, target); got != want {
			t.Errorf("lowerBound(%d) = %d, want %d", target, got, want)
		}
	}
}
