// functions89
// Make the tests pass!

// I AM NOT DONE
//
// search(nums, target, lo, hi) looks for target in the sorted slice over [lo, hi).
// It returns the index, or -1.
// Practices recursion with a shrinking range.
package main_test

import "testing"

func search(nums []int, target, lo, hi int) int {
	if lo >= hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	switch {
	case nums[mid] == target:
		return mid
	case target < nums[mid]:
		return search(nums, target, mid, hi)
	default:
		return search(nums, target, lo, mid)
	}
}

func TestSearch(t *testing.T) {
	nums := []int{1, 4, 7, 9, 12, 20}
	for i, v := range nums {
		if got := search(nums, v, 0, len(nums)); got != i {
			t.Errorf("search(%d) = %d, want %d", v, got, i)
		}
	}
	if got := search(nums, 5, 0, len(nums)); got != -1 {
		t.Errorf("search(5) = %d, want -1", got)
	}
}
