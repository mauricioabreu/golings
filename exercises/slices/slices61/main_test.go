// slices61
// Make the tests pass!

// I AM NOT DONE
//
// clone copies the data, yet it returns an empty slice.
// copy moves min(len(dst), len(src)) elements; the capacity does not count.
package main_test

import (
	"reflect"
	"testing"
)

func clone(s []int) []int {
	out := make([]int, 0, len(s))
	copy(out, s)
	return out
}

func TestClone(t *testing.T) {
	src := []int{1, 2, 3}
	c := clone(src)
	src[0] = 9
	if !reflect.DeepEqual(c, []int{1, 2, 3}) {
		t.Errorf("clone = %v", c)
	}
}
