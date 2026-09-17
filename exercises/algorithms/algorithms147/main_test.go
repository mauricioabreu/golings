// algorithms147
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: binary search for the lower bound. Return the index of target
// in a sorted slice, or the position where it should be inserted.
// Expected asymptotics: O(log n) time, O(1) space.
package main_test

import "testing"

func searchInsert(nums []int, target int) int {
	return 0
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	cases := map[int]int{5: 2, 2: 1, 7: 4, 0: 0, 6: 3}
	for target, want := range cases {
		if got := searchInsert(nums, target); got != want {
			t.Errorf("searchInsert(%d) = %d, want %d", target, got, want)
		}
	}
	if searchInsert(nil, 5) != 0 {
		t.Errorf("searchInsert(nil) should be 0")
	}
}
