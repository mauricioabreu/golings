// slices33
// Make the tests pass!

// I AM NOT DONE
//
// desc sorts the numbers in descending order.
// Practices sort.Sort, sort.Reverse and sort.IntSlice.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func desc(s []int) {
	sort.Sort(sort.IntSlice(s))
}

func TestDesc(t *testing.T) {
	s := []int{2, 9, 4}
	desc(s)
	if !reflect.DeepEqual(s, []int{9, 4, 2}) {
		t.Errorf("desc = %v", s)
	}
}
