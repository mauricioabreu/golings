// generics91
// Make the tests pass!

// I AM NOT DONE
//
// Optional[T] holds a value along with a flag saying whether it is there.
// Practices generic structs with a constructor.
package main_test

import "testing"

type Optional[T any] struct {
	value T
	ok    bool
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{value: v}
}

func (o Optional[T]) IsSet() bool { return o.ok }

func TestOptional(t *testing.T) {
	if !Some(0).IsSet() || (Optional[string]{}).IsSet() {
		t.Errorf("Optional works incorrectly")
	}
}
