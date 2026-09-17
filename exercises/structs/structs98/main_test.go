// structs98
// Make the tests pass!

// I AM NOT DONE
//
// Area returns the area of a rectangle.
// Practices methods with a value receiver.
package main_test

import "testing"

type Rect struct{ W, H float64 }

func (r Rect) Area() float64 {
	return r.W + r.H
}

func TestArea(t *testing.T) {
	if got := (Rect{3, 4}).Area(); got != 12 {
		t.Errorf("Area = %v", got)
	}
}
