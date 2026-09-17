// arrays23
// Make the tests pass!

// I AM NOT DONE
//
// sameArgs compares two argument sets of type [2]interface{}.
// When the arguments hold slices, a comparison with == panics.
// An array of interfaces is comparable at compile time, yet it may panic.
package main_test

import (
	"reflect"
	"testing"
)

func sameArgs(a, b [2]interface{}) bool {
	return a == b
}

func TestSameArgs(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameArgs([2]interface{}{1, "x"}, [2]interface{}{1, "x"}) {
		t.Errorf("scalars should match")
	}
	if !sameArgs([2]interface{}{[]int{1}, 2}, [2]interface{}{[]int{1}, 2}) {
		t.Errorf("equal slices should match")
	}
}
