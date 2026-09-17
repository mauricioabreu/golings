// generics98
// Make the tests pass!

// I AM NOT DONE
//
// IsNil has to return true for a nil pointer of any type as well.
// any(v) == nil does not match for a nil pointer.
// An interface holding a typed nil is not nil.
package main_test

import (
	"reflect"
	"testing"
)

func IsNil[T any](v T) bool {
	return any(v) == nil
}

func TestIsNil(t *testing.T) {
	_ = reflect.ValueOf
	var p *int
	var e error
	if !IsNil(p) || !IsNil(e) || IsNil(5) || IsNil(new(int)) {
		t.Errorf("IsNil works incorrectly")
	}
}
