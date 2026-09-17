// switch10
// Make the tests pass!

// I AM NOT DONE
//
// unitFactor returns the factor for the units "cm", "m" and "km".
// The code does not compile: a switch without a default is not a terminating statement.
// Practices the terminating statement rules of Go.
package main_test

import "testing"

func unitFactor(unit string) int {
	switch unit {
	case "cm":
		return 1
	case "m":
		return 100
	case "km":
		return 100000
	}
}

func TestUnitFactor(t *testing.T) {
	cases := map[string]int{"cm": 1, "m": 100, "km": 100000, "mile": 0}
	for in, want := range cases {
		if got := unitFactor(in); got != want {
			t.Errorf("unitFactor(%s) = %d, want %d", in, got, want)
		}
	}
}
