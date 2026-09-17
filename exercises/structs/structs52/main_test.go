// structs52
// Make the tests pass!

// I AM NOT DONE
//
// Copies of Widget share a *Style, and changing the colour of a copy changes the original.
// Copying a struct copies the pointer, not the object.
package main_test

import "testing"

type Style struct{ Color string }

type Widget struct {
	*Style
	Label string
}

func (w Widget) WithColor(c string) Widget {
	w.Color = c
	return w
}

func TestWithColor(t *testing.T) {
	base := Widget{&Style{"black"}, "ok"}
	red := base.WithColor("red")
	if red.Color != "red" || base.Color != "black" {
		t.Errorf("red = %s, base = %s", red.Color, base.Color)
	}
}
