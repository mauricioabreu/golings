// algorithms102
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: insertion sort. Sort a slice in ascending order in place,
// without using the sort package.
// Expected asymptotics: O(n²) time, O(1) space.
package main_test

import (
	"reflect"
	"testing"
)

func insertionSort(nums []int) {
}

func TestInsertionSort(t *testing.T) {
	cases := [][]int{
		{5, 2, 4, 6, 1, 3},
		{},
		{1},
		{3, 3, 3},
		{5, 4, 3, 2, 1},
		{-2, 0, -5, 7},
	}
	for _, c := range cases {
		got := append([]int{}, c...)
		insertionSort(got)
		want := append([]int{}, c...)
		for i := 1; i < len(want); i++ {
			for j := i; j > 0 && want[j] < want[j-1]; j-- {
				want[j], want[j-1] = want[j-1], want[j]
			}
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("insertionSort(%v) = %v, want %v", c, got, want)
		}
	}
}
