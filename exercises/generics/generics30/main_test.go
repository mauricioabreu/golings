// generics30
// Make the tests pass!

// I AM NOT DONE
//
// Box[T] holds a value, and Get returns it.
// Practices methods on a generic type.
package main_test

import "testing"

type Box[T any] struct{ v T }

func (b Box[T]) Get() T {
	return *new(T)
}

func TestBox(t *testing.T) {
	if (Box[string]{"x"}).Get() != "x" || (Box[int]{7}).Get() != 7 {
		t.Errorf("Box works incorrectly")
	}
}
