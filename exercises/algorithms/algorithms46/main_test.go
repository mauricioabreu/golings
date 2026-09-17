// algorithms46
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: modified binary search. A sorted slice of distinct numbers has
// been rotated. Return the index of target, or -1.
// Expected asymptotics: O(log n) time, O(1) space.
package main_test

import "testing"

func searchRotated(nums []int, target int) int {
	return 0
}

func TestSearchRotated(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	cases := map[int]int{0: 4, 3: -1, 4: 0, 2: 6, 7: 3}
	for target, want := range cases {
		if got := searchRotated(nums, target); got != want {
			t.Errorf("searchRotated(%d) = %d, want %d", target, got, want)
		}
	}
	if searchRotated([]int{1}, 0) != -1 || searchRotated(nil, 1) != -1 || searchRotated([]int{3, 1}, 1) != 1 {
		t.Errorf("edge cases failed")
	}
}
