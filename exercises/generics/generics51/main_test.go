// generics51
// Make the tests pass!

// I AM NOT DONE
//
// MapOpt transforms an Optional value while keeping the absence of a value.
// Practices generic functions over generic types.
package main_test

import (
	"strconv"
	"testing"
)

type Optional[T any] struct {
	Val T
	OK  bool
}

func MapOpt[T, U any](o Optional[T], f func(T) U) Optional[U] {
	return Optional[U]{Val: f(o.Val), OK: true}
}

func TestMapOpt(t *testing.T) {
	some := MapOpt(Optional[int]{7, true}, strconv.Itoa)
	none := MapOpt(Optional[int]{}, strconv.Itoa)
	if some.Val != "7" || !some.OK || none.OK || none.Val != "" {
		t.Errorf("some=%+v none=%+v", some, none)
	}
}
