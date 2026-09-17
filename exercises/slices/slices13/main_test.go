// slices13
// Make the tests pass!

// I AM NOT DONE
//
// trimZeros strips the zeros from both ends of a slice.
// Practices moving the bounds of a slice.
package main_test

import (
	"reflect"
	"testing"
)

func trimZeros(s []int) []int {
	start, end := 0, len(s)
	for s[start] == 0 {
		start++
	}
	return s[start:end]
}

func TestTrimZeros(t *testing.T) {
	if got := trimZeros([]int{0, 0, 1, 0, 2, 0}); !reflect.DeepEqual(got, []int{1, 0, 2}) {
		t.Errorf("trimZeros = %v", got)
	}
	if got := trimZeros([]int{0, 0}); len(got) != 0 {
		t.Errorf("trimZeros(all zeros) = %v", got)
	}
}
