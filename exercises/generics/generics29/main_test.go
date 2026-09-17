// generics29
// Make the tests pass!

// I AM NOT DONE
//
// Box[int] and Box[int64] are different types. The assignment does not compile.
// Instantiations of a generic type are not compatible with each other.
package main_test

import "testing"

type Box[T any] struct{ V T }

func widen(b Box[int]) Box[int64] {
	return b
}

func TestWiden(t *testing.T) {
	if widen(Box[int]{7}).V != 7 {
		t.Errorf("widen failed")
	}
}
