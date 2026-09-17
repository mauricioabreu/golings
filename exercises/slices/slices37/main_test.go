// slices37
// Make the tests pass!

// I AM NOT DONE
//
// sortedCopy returns a sorted copy.
// Practices sort.Ints.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func sortedCopy(s []int) []int {
	c := append([]int(nil), s...)
	sort.Ints(c)
	return s
}

func TestSortedCopy(t *testing.T) {
	in := []int{3, 1, 2}
	if got := sortedCopy(in); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("sortedCopy = %v", got)
	}
}
