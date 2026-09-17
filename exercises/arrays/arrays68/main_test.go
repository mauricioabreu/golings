// arrays68
// Make the tests pass!

// I AM NOT DONE
//
// samePair compares pairs of lists. The code does not compile: an array of slices
// cannot be compared with ==.
// An array is comparable only when its element type is.
package main_test

import (
	"reflect"
	"testing"
)

func samePair(a, b [2][]int) bool {
	return a == b
}

func TestSamePair(t *testing.T) {
	_ = reflect.DeepEqual
	a := [2][]int{{1, 2}, {3}}
	b := [2][]int{{1, 2}, {3}}
	c := [2][]int{{1}, {3}}
	if !samePair(a, b) || samePair(a, c) {
		t.Errorf("samePair works incorrectly")
	}
}
