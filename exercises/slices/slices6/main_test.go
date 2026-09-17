// slices6
// Make the tests pass!

// I AM NOT DONE
//
// shiftRight moves the elements one place right, dropping the last one,
// and leaves a zero in front. Right now the whole slice is filled with the first value.
// Practices the copying order when the ranges overlap.
package main_test

import (
	"reflect"
	"testing"
)

func shiftRight(s []int) {
	for i := 1; i < len(s); i++ {
		s[i] = s[i-1]
	}
	if len(s) > 0 {
		s[0] = 0
	}
}

func TestShiftRight(t *testing.T) {
	s := []int{1, 2, 3, 4}
	shiftRight(s)
	if !reflect.DeepEqual(s, []int{0, 1, 2, 3}) {
		t.Errorf("shiftRight = %v", s)
	}
}
