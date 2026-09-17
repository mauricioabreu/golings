// slices32
// Make the tests pass!

// I AM NOT DONE
//
// removeFirst removes the first element equal to x. The caller keeps using
// the old slice and sees a duplicate left at the end.
// A function can change the elements, but not the length of the caller's slice.
package main_test

import (
	"reflect"
	"testing"
)

func removeFirst(s []int, x int) {
	for i, v := range s {
		if v == x {
			s = append(s[:i], s[i+1:]...)
			return
		}
	}
}

func cleanup(s []int) []int {
	removeFirst(s, 0)
	return s
}

func TestCleanup(t *testing.T) {
	if got := cleanup([]int{1, 0, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("cleanup = %v, want [1 2 3]", got)
	}
}
