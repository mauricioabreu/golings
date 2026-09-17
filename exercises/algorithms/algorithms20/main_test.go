// algorithms20
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers. The slice is sorted and may contain negative
// numbers. Return the sorted squares of the elements without sorting again.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import (
	"reflect"
	"testing"
)

func sortedSquares(nums []int) []int {
	return nil
}

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{}, []int{}},
		{[]int{-5}, []int{25}},
		{[]int{-3, -2, -1}, []int{1, 4, 9}},
	}
	for _, c := range cases {
		if got := sortedSquares(c.nums); !reflect.DeepEqual(got, c.want) {
			t.Errorf("sortedSquares(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
