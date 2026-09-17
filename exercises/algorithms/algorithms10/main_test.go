// algorithms10
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a monotonic stack. For each element, return the first element
// to its right that is greater; if there is none — -1.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import (
	"reflect"
	"testing"
)

func nextGreater(nums []int) []int {
	return nil
}

func TestNextGreater(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{2, 1, 2, 4, 3}, []int{4, 2, 4, -1, -1}},
		{[]int{5, 4, 3}, []int{-1, -1, -1}},
		{[]int{1}, []int{-1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := nextGreater(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("nextGreater(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
