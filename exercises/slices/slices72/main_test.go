// slices72
// Make the tests pass!

// I AM NOT DONE
//
// insert puts x at position i. The result holds duplicates instead of a shift.
// append(append(s[:i], x), s[i:]...) overwrites s[i] before the tail is copied.
package main_test

import (
	"reflect"
	"testing"
)

func insert(s []int, i, x int) []int {
	return append(append(s[:i], x), s[i:]...)
}

func TestInsert(t *testing.T) {
	s := make([]int, 3, 10)
	copy(s, []int{1, 2, 4})
	if got := insert(s, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("insert = %v, want [1 2 3 4]", got)
	}
}
