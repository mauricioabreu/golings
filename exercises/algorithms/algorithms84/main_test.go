// algorithms84
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a sliding window with an ordered structure. For each window
// of size k return the median (for even k — the mean of the two middle values).
// Expected asymptotics: O(n·k) time (or O(n log k) with heaps), O(k) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	return nil
}

func TestMedianSlidingWindow(t *testing.T) {
	_ = sort.Ints
	cases := []struct {
		nums []int
		k    int
		want []float64
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1, -1, -1, 3, 5, 6}},
		{[]int{1, 2, 3, 4}, 2, []float64{1.5, 2.5, 3.5}},
		{[]int{5}, 1, []float64{5}},
		{[]int{2147483647, 2147483647}, 2, []float64{2147483647}},
		{[]int{1}, 2, nil},
	}
	for _, c := range cases {
		if got := medianSlidingWindow(c.nums, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("medianSlidingWindow(%v, %d) = %v, want %v", c.nums, c.k, got, c.want)
		}
	}
}
