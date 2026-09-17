// switch96
// Make the tests pass!

// I AM NOT DONE
//
// shirtSize: up to 90 cm is "S", up to 100 "M", up to 110 "L", anything else "XL".
// The branches overlap, and the first one swallows the rest.
// The first matching branch of a switch is the one that runs.
package main_test

import "testing"

func shirtSize(chest int) string {
	switch {
	case chest < 110:
		return "L"
	case chest < 100:
		return "M"
	case chest < 90:
		return "S"
	}
	return "XL"
}

func TestShirtSize(t *testing.T) {
	cases := map[int]string{85: "S", 90: "M", 99: "M", 105: "L", 110: "XL"}
	for in, want := range cases {
		if got := shirtSize(in); got != want {
			t.Errorf("shirtSize(%d) = %s, want %s", in, got, want)
		}
	}
}
