// anonymous_functions24
// Make the tests pass!

// I AM NOT DONE
//
// firstAtLeast returns the index of the first value >= x in a sorted slice.
// Practices a literal for sort.Search.
package main_test

import (
	"sort"
	"testing"
)

func firstAtLeast(s []int, x int) int {
	return sort.Search(len(s), func(i int) bool {
		return s[i] <= x
	})
}

func TestFirstAtLeast(t *testing.T) {
	s := []int{1, 4, 4, 9}
	if firstAtLeast(s, 4) != 1 || firstAtLeast(s, 5) != 3 || firstAtLeast(s, 10) != 4 {
		t.Errorf("firstAtLeast works incorrectly")
	}
}
