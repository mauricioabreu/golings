// slices30
// Make the tests pass!

// I AM NOT DONE
//
// interleave alternates the elements of two slices: a0, b0, a1, b1, ...,
// with the tail of the longer slice appended at the end.
// Practices one index over two slices.
package main_test

import (
	"reflect"
	"testing"
)

func interleave(a, b []int) []int {
	var out []int
	for i := 0; i < len(a); i++ {
		out = append(out, a[i], b[i])
	}
	return out
}

func TestInterleave(t *testing.T) {
	if got := interleave([]int{1, 3, 5, 7}, []int{2, 4}); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 7}) {
		t.Errorf("interleave = %v", got)
	}
	if got := interleave(nil, []int{9}); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("interleave(nil, [9]) = %v", got)
	}
}
