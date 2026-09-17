// slices48
// Make the tests pass!

// I AM NOT DONE
//
// insertAt inserts an element at position i, shifting the rest to the right.
// Practices append to grow and copy to shift.
package main_test

import (
	"reflect"
	"testing"
)

func insertAt(s []int, i, x int) []int {
	s = append(s, 0)
	copy(s[i:], s[i+1:])
	s[i+1] = x
	return s
}

func TestInsertAt(t *testing.T) {
	if got := insertAt([]int{1, 2, 4}, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("insertAt = %v", got)
	}
	if got := insertAt([]int{2}, 0, 1); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("insertAt(0) = %v", got)
	}
	if got := insertAt([]int{1}, 1, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("insertAt(end) = %v", got)
	}
}
