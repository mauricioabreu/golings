// switch47
// Make the tests pass!

// I AM NOT DONE
//
// planetOrder returns the position of a planet from the Sun.
// Practices a switch on a string.
package main_test

import "testing"

func planetOrder(p string) int {
	switch p {
	case "Mercury":
		return 1
	case "Venus":
		return 2
	case "Earth":
		return 4
	case "Mars":
		return 4
	}
	return 0
}

func TestPlanetOrder(t *testing.T) {
	cases := map[string]int{"Mercury": 1, "Venus": 2, "Earth": 3, "Mars": 4, "Pluto": 0}
	for in, want := range cases {
		if got := planetOrder(in); got != want {
			t.Errorf("planetOrder(%s) = %d, want %d", in, got, want)
		}
	}
}
