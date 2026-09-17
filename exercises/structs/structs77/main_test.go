// structs77
// Make the tests pass!

// I AM NOT DONE
//
// Rect holds a width and a height; Perimeter, and Scale which scales in place.
// Practices mixing value and pointer receivers.
package main_test

import "testing"

type Rect struct{ W, H float64 }

func (r Rect) Perimeter() float64 { return r.W + r.H }
func (r Rect) Scale(k float64)    { r.W *= k; r.H *= k }

func TestRect(t *testing.T) {
	r := Rect{2, 3}
	r.Scale(2)
	if r.W != 4 || r.H != 6 || r.Perimeter() != 20 {
		t.Errorf("rect = %+v, perimeter %v", r, r.Perimeter())
	}
}
