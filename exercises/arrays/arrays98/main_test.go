// arrays98
// Make the tests pass!

// I AM NOT DONE
//
// extend appends an element to a slice of the whole array and then changes the first element.
// The changes are expected to land in the original array, and they do not.
// append allocates a new array when the capacity runs out.
package main_test

import (
	"reflect"
	"testing"
)

func extend(a *[3]int, v int) []int {
	s := a[:]
	s = append(s, v)
	s[0] = -1
	return s
}

func TestExtend(t *testing.T) {
	a := [3]int{1, 2, 3}
	s := extend(&a, 4)
	if a[0] != -1 {
		t.Errorf("array not updated: %v", a)
	}
	if !reflect.DeepEqual(s, []int{-1, 2, 3, 4}) {
		t.Errorf("slice = %v", s)
	}
}
