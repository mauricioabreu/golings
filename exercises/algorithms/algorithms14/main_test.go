// algorithms14
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: prefixes and suffixes. For each position i, return the product of
// all elements except nums[i], without using division.
// Expected asymptotics: O(n) time, O(1) extra space (excluding the answer).
package main_test

import (
	"reflect"
	"testing"
)

func productExceptSelf(nums []int) []int {
	return nil
}

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{5, 7}, []int{7, 5}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := productExceptSelf(c.nums); !reflect.DeepEqual(got, c.want) {
			t.Errorf("productExceptSelf(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
