// algorithms30
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sorting and two pointers. Find all unique triples of numbers
// with a sum of 0. Each triple is ordered ascending; the triples are
// ordered lexicographically.
// Expected asymptotics: O(n²) time, O(1) extra space (ignoring sorting
// and the answer).
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func threeSum(nums []int) [][3]int {
	return nil
}

func TestThreeSum(t *testing.T) {
	_ = sort.Ints
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	if !reflect.DeepEqual(got, [][3]int{{-1, -1, 2}, {-1, 0, 1}}) {
		t.Errorf("threeSum = %v", got)
	}
	if got := threeSum([]int{0, 0, 0, 0}); !reflect.DeepEqual(got, [][3]int{{0, 0, 0}}) {
		t.Errorf("threeSum(zeros) = %v", got)
	}
	if got := threeSum([]int{0, 1, 1}); len(got) != 0 {
		t.Errorf("threeSum(no triples) = %v", got)
	}
	if got := threeSum(nil); len(got) != 0 {
		t.Errorf("threeSum(nil) = %v", got)
	}
}
