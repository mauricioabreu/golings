// algorithms105
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a binary heap in an array. Sort a slice in place, without
// using the sort and container/heap packages.
// Expected asymptotics: O(n·log n) time, O(1) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func heapSort(nums []int) {
}

func TestHeapSort(t *testing.T) {
	_ = sort.Ints
	cases := [][]int{
		{4, 10, 3, 5, 1},
		{},
		{1},
		{2, 2, 1},
		{-1, -5, 3, 0},
	}
	for _, c := range cases {
		got := append([]int{}, c...)
		heapSort(got)
		want := append([]int{}, c...)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("heapSort(%v) = %v, want %v", c, got, want)
		}
	}
	big := make([]int, 10000)
	for i := range big {
		big[i] = (i * 31) % 9973
	}
	heapSort(big)
	if !sort.IntsAreSorted(big) {
		t.Errorf("big slice is not sorted")
	}
}
