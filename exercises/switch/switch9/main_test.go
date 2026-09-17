// switch9
// Make the tests pass!

// I AM NOT DONE
//
// hexColor returns the HEX code of a primary colour.
// Practices a switch on a string.
package main_test

import "testing"

func hexColor(name string) string {
	switch name {
	case "red":
		return "#FF0000"
	case "green":
		return "#0000FF"
	case "blue":
		return "#0000FF"
	}
	return "#000000"
}

func TestHexColor(t *testing.T) {
	cases := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF", "pink": "#000000"}
	for in, want := range cases {
		if got := hexColor(in); got != want {
			t.Errorf("hexColor(%s) = %s, want %s", in, got, want)
		}
	}
}
