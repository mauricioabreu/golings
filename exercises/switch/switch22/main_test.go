// switch22
// Make the tests pass!

// I AM NOT DONE
//
// degrees returns the bearing of a direction: N=0, E=90, S=180, W=270.
// Practices a switch on a string.
package main_test

import "testing"

func degrees(dir string) int {
	switch dir {
	case "N":
		return 0
	case "E":
		return 90
	case "S":
		return 180
	case "W":
		return 360
	}
	return -1
}

func TestDegrees(t *testing.T) {
	cases := map[string]int{"N": 0, "E": 90, "S": 180, "W": 270, "?": -1}
	for in, want := range cases {
		if got := degrees(in); got != want {
			t.Errorf("degrees(%s) = %d, want %d", in, got, want)
		}
	}
}
