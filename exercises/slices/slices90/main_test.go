// slices90
// Make the tests pass!

// I AM NOT DONE
//
// prepend puts an element at the front of a slice.
// Practices appending to a new slice and spreading.
package main_test

import (
	"reflect"
	"testing"
)

func prepend(s []int, x int) []int {
	return append(s, x)
}

func TestPrepend(t *testing.T) {
	if got := prepend([]int{2, 3}, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("prepend = %v", got)
	}
}
