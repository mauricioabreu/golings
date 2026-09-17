// algorithms144
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers. In a sorted slice, keep only unique values in
// place at the beginning and return their count.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import (
	"reflect"
	"testing"
)

func removeDuplicates(nums []int) int {
	return 0
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{nil, []int{}},
		{[]int{7}, []int{7}},
		{[]int{-3, -3, -3}, []int{-3}},
	}
	for _, c := range cases {
		in := append([]int{}, c.nums...)
		n := removeDuplicates(in)
		if !reflect.DeepEqual(in[:n], c.want) {
			t.Errorf("removeDuplicates(%v) -> %v, want %v", c.nums, in[:n], c.want)
		}
	}
}
