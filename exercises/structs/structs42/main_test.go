// structs42
// Make the tests pass!

// I AM NOT DONE
//
// newColor builds a colour with a positional literal, the fields being R, G, B.
// Practices a struct literal without field names.
package main_test

import "testing"

type Color struct{ R, G, B uint8 }

func orange() Color {
	return Color{0, 165, 255}
}

func TestOrange(t *testing.T) {
	if c := orange(); c.R != 255 || c.G != 165 || c.B != 0 {
		t.Errorf("orange = %+v", c)
	}
}
