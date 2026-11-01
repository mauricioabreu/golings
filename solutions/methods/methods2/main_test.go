// methods2
// Make the tests pass!

package main_test

import "testing"

type Rectangle struct {
	Width, Height int
}

// Scale multiplies both sides of the rectangle by factor.
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

func TestScale(t *testing.T) {
	r := Rectangle{Width: 2, Height: 3}

	r.Scale(2)

	if r.Width != 4 || r.Height != 6 {
		t.Errorf("got %dx%d, want 4x6", r.Width, r.Height)
	}
}
