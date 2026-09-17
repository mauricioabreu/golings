// slices66
// Make the tests pass!

// I AM NOT DONE
//
// insertSorted inserts a number into a sorted slice, keeping the order.
// Practices sort.SearchInts and inserting at an index.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func insertSorted(s []int, x int) []int {
	i := sort.SearchInts(s, x)
	s = append(s, x)
	_ = i
	return s
}

func TestInsertSorted(t *testing.T) {
	s := []int{}
	for _, v := range []int{5, 1, 3, 4, 2} {
		s = insertSorted(s, v)
	}
	if !reflect.DeepEqual(s, []int{1, 2, 3, 4, 5}) {
		t.Errorf("insertSorted = %v", s)
	}
}
