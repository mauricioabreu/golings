// slices28
// Make the tests pass!

// I AM NOT DONE
//
// intersect returns the elements that occur in both slices, with no repeats
// and in the order of the first slice.
// Practices building a set out of a slice.
package main_test

import (
	"reflect"
	"testing"
)

func intersect(a, b []int) []int {
	set := map[int]bool{}
	for _, v := range a {
		set[v] = true
	}
	var out []int
	for _, v := range b {
		if set[v] {
			out = append(out, v)
		}
	}
	return out
}

func TestIntersect(t *testing.T) {
	if got := intersect([]int{4, 1, 4, 2, 3}, []int{3, 4, 5, 4}); !reflect.DeepEqual(got, []int{4, 3}) {
		t.Errorf("intersect = %v, want [4 3]", got)
	}
}
