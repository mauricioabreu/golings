// slices46
// Make the tests pass!

// I AM NOT DONE
//
// reversed returns a new slice with the elements in reverse order.
// Practices computing an index from the end.
package main_test

import (
	"reflect"
	"testing"
)

func reversed(s []int) []int {
	out := make([]int, len(s))
	for i := range s {
		out[i] = s[len(s)-i]
	}
	return out
}

func TestReversed(t *testing.T) {
	if got := reversed([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("reversed = %v", got)
	}
}
