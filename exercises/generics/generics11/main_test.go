// generics11
// Make the tests pass!

// I AM NOT DONE
//
// FirstOr returns the first element, or a "nothing" value. The code does not compile:
// nil cannot be used as a value of a type parameter.
// The zero value of T comes from var zero T.
package main_test

import "testing"

func FirstOr[T any](s []T) (T, bool) {
	if len(s) == 0 {
		return nil, false
	}
	return s[0], true
}

func TestFirstOr(t *testing.T) {
	if v, ok := FirstOr([]int{}); ok || v != 0 {
		t.Errorf("FirstOr(empty) = %v, %v", v, ok)
	}
	if v, ok := FirstOr([]string{"x"}); !ok || v != "x" {
		t.Errorf("FirstOr = %v, %v", v, ok)
	}
}
