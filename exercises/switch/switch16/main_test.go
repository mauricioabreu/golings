// switch16
// Make the tests pass!

// I AM NOT DONE
//
// sides returns the number of sides of a shape.
// Practices a switch on a string returning a value.
package main_test

import "testing"

func sides(shape string) int {
	switch shape {
	case "triangle":
		return 3
	case "square":
		return 3
	case "hexagon":
		return 6
	}
	return 0
}

func TestSides(t *testing.T) {
	cases := map[string]int{"triangle": 3, "square": 4, "hexagon": 6, "circle": 0}
	for in, want := range cases {
		if got := sides(in); got != want {
			t.Errorf("sides(%s) = %d, want %d", in, got, want)
		}
	}
}
