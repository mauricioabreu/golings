// slices11
// Make the tests pass!

// I AM NOT DONE
//
// merge combines two sorted slices into one sorted slice.
// Practices two indexes and appending the tails.
package main_test

import (
	"reflect"
	"testing"
)

func merge(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	return out
}

func TestMerge(t *testing.T) {
	if got := merge([]int{1, 4, 7}, []int{2, 3, 8, 9}); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 7, 8, 9}) {
		t.Errorf("merge = %v", got)
	}
	if got := merge(nil, []int{1}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("merge(nil, [1]) = %v", got)
	}
}
