// generics57
// Make the tests pass!

// I AM NOT DONE
//
// Double doubles values, the named type Celsius included.
// The code does not compile: the constraint only accepts float64 itself.
// ~T covers every type whose underlying type is T.
package main_test

import "testing"

type Celsius float64

type Float interface {
	float64
}

func Double[T Float](v T) T { return v * 2 }

func TestDouble(t *testing.T) {
	if Double(Celsius(21.5)) != 43 {
		t.Errorf("Double = %v", Double(Celsius(21.5)))
	}
}
