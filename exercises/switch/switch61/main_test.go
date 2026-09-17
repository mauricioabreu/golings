// switch61
// Make the tests pass!

// I AM NOT DONE
//
// opposite returns the opposite direction: N<->S, E<->W.
// Practices a switch on string values.
package main_test

import "testing"

func opposite(d string) string {
	switch d {
	case "N":
		return "S"
	case "S":
		return "N"
	case "E":
		return "E"
	case "W":
		return "E"
	}
	return ""
}

func TestOpposite(t *testing.T) {
	cases := map[string]string{"N": "S", "S": "N", "E": "W", "W": "E", "X": ""}
	for in, want := range cases {
		if got := opposite(in); got != want {
			t.Errorf("opposite(%s) = %q, want %q", in, got, want)
		}
	}
}
