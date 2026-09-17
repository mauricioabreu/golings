// functions98
// Make the tests pass!

// I AM NOT DONE
//
// sortDesc must sort the numbers in descending order.
// Practices passing a comparison function to sort.Slice.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func sortDesc(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

func TestSortDesc(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5}
	sortDesc(nums)
	if !reflect.DeepEqual(nums, []int{5, 4, 3, 1, 1}) {
		t.Errorf("sortDesc = %v, want [5 4 3 1 1]", nums)
	}
}
