// generics40
// Make the tests pass!

// I AM NOT DONE
//
// Pair{1, "a"} does not compile: type arguments are not inferred for
// composite literals.
// An instantiation is written out explicitly, or comes from a constructor function.
package main_test

import "testing"

type Pair[A, B any] struct {
	First  A
	Second B
}

func sample() Pair[int, string] {
	return Pair{1, "a"}
}

func TestSample(t *testing.T) {
	if p := sample(); p.First != 1 || p.Second != "a" {
		t.Errorf("sample = %+v", p)
	}
}
