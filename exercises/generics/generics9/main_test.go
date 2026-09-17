// generics9
// Make the tests pass!

// I AM NOT DONE
//
// As tries to convert a value to the type T. The code does not compile:
// a type assertion only works on an interface value.
// v.(T) needs v to be an interface.
package main_test

import "testing"

func As[T any](v T) (T, bool) {
	r, ok := v.(T)
	return r, ok
}

func TestAs(t *testing.T) {
	if s, ok := As[string]("go"); !ok || s != "go" {
		t.Errorf("As[string] = %q, %v", s, ok)
	}
	if _, ok := As[int]("go"); ok {
		t.Errorf("As[int](\"go\") should fail")
	}
}
