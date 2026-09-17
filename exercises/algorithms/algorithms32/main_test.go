// algorithms32
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: divide and conquer with Lomuto partitioning. Sort a slice
// in place without using the sort package.
// Expected asymptotics: O(n·log n) on average, O(log n) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func quickSort(nums []int) {
}

func TestQuickSort(t *testing.T) {
	_ = sort.Ints
	cases := [][]int{
		{3, 6, 1, 8, 2, 9, 4},
		{},
		{1},
		{2, 1},
		{5, 5, 5, 5},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
	for _, c := range cases {
		got := append([]int{}, c...)
		quickSort(got)
		want := append([]int{}, c...)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("quickSort(%v) = %v, want %v", c, got, want)
		}
	}
	big := make([]int, 20000)
	for i := range big {
		big[i] = (i * 7919) % 20011
	}
	quickSort(big)
	if !sort.IntsAreSorted(big) {
		t.Errorf("big slice is not sorted")
	}
}
