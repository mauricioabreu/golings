// algorithms25
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers. The slice is sorted in non-decreasing order.
// Return the indices i < j of a pair with the sum target, or (-1, -1).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func twoSumSorted(nums []int, target int) (int, int) {
	return 0, 0
}

func TestTwoSumSorted(t *testing.T) {
	cases := []struct {
		nums         []int
		target, i, j int
	}{
		{[]int{2, 7, 11, 15}, 9, 0, 1},
		{[]int{2, 3, 4}, 6, 0, 2},
		{[]int{-1, 0}, -1, 0, 1},
		{[]int{1, 2, 3}, 7, -1, -1},
		{nil, 1, -1, -1},
		{[]int{4}, 8, -1, -1},
	}
	for _, c := range cases {
		if i, j := twoSumSorted(c.nums, c.target); i != c.i || j != c.j {
			t.Errorf("twoSumSorted(%v, %d) = (%d, %d), want (%d, %d)", c.nums, c.target, i, j, c.i, c.j)
		}
	}
}
