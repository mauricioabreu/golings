// slices12
// Make the tests pass!

// I AM NOT DONE
//
// indexOf looks a number up in a sorted slice with sort.Search and returns -1
// when it is not there. For missing numbers it returns somebody else's index.
// sort.Search returns an insertion point, not a found flag.
package main_test

import (
	"sort"
	"testing"
)

func indexOf(s []int, x int) int {
	i := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if i == len(s) {
		return -1
	}
	return i
}

func TestIndexOf(t *testing.T) {
	s := []int{2, 4, 6, 8}
	if indexOf(s, 6) != 2 || indexOf(s, 2) != 0 {
		t.Errorf("indexOf of present values failed")
	}
	if indexOf(s, 5) != -1 || indexOf(s, 9) != -1 || indexOf(s, 1) != -1 {
		t.Errorf("indexOf of missing values should be -1")
	}
}
