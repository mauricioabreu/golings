// generics64
// Make the tests pass!

// I AM NOT DONE
//
// Negate returns -v, and the constraint lists unsigned types.
// The code does not compile.
// -1 cannot be represented in a uint, so signed arithmetic calls for signed types.
package main_test

import "testing"

type Signed interface {
	~int | ~int64 | ~uint
}

func Negate[T Signed](v T) T {
	return v * -1
}

func TestNegate(t *testing.T) {
	if Negate(5) != -5 || Negate(int64(-3)) != 3 {
		t.Errorf("Negate works incorrectly")
	}
}
