// algorithms86
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: three pointers. The slice contains only 0, 1, and 2. Sort it
// in place in a single pass without counting.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import (
	"reflect"
	"testing"
)

func sortColors(nums []int) {
}

func TestSortColors(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 2, 0, 0}, []int{0, 0, 2, 2}},
	}
	for _, c := range cases {
		got := append([]int{}, c.nums...)
		sortColors(got)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("sortColors(%v) = %v", c.nums, got)
		}
	}
}
