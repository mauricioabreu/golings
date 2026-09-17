// structs76
// Make the tests pass!

// I AM NOT DONE
//
// grow raises the radius of a circle held in an interface.
// The code does not compile: the result of a type assertion is not addressable.
// A value inside an interface cannot be changed in place.
package main_test

import "testing"

type Circle struct{ R float64 }

func grow(s interface{}) {
	s.(Circle).R *= 2
}

func TestGrow(t *testing.T) {
	c := &Circle{R: 1.5}
	grow(c)
	if c.R != 3 {
		t.Errorf("R = %v, want 3", c.R)
	}
}
