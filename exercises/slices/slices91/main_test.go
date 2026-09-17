// slices91
// Make the tests pass!

// I AM NOT DONE
//
// rotateLeft returns a new slice shifted k places to the left.
// Practices joining two parts of a slice.
package main_test

import (
	"reflect"
	"testing"
)

func rotateLeft(s []int, k int) []int {
	return append(s[:k], s[k:]...)
}

func TestRotateLeft(t *testing.T) {
	if got := rotateLeft([]int{1, 2, 3, 4, 5}, 2); !reflect.DeepEqual(got, []int{3, 4, 5, 1, 2}) {
		t.Errorf("rotateLeft(2) = %v", got)
	}
	if got := rotateLeft([]int{1, 2, 3}, 4); !reflect.DeepEqual(got, []int{2, 3, 1}) {
		t.Errorf("rotateLeft(4) = %v", got)
	}
	if got := rotateLeft(nil, 1); len(got) != 0 {
		t.Errorf("rotateLeft(nil) = %v", got)
	}
}
