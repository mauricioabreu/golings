// if25
// Make the tests pass!

// I AM NOT DONE
//
// sameValue compares two interface{} values. For slices a comparison
// with == panics at run time.
// == on interfaces holding an incomparable dynamic type panics.
package main_test

import (
	"reflect"
	"testing"
)

func sameValue(a, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

func TestSameValue(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameValue(1, 1) || sameValue(1, "1") {
		t.Errorf("scalar comparison is wrong")
	}
	if !sameValue([]int{1, 2}, []int{1, 2}) {
		t.Errorf("equal slices should be the same value")
	}
	if sameValue([]int{1}, []int{2}) {
		t.Errorf("different slices should not be the same value")
	}
}
