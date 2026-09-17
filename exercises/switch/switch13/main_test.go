// switch13
// Make the tests pass!

// I AM NOT DONE
//
// cents turns the name of a coin into cents.
// Practices a switch that returns a value.
package main_test

import "testing"

func cents(coin string) int {
	switch coin {
	case "penny":
		return 1
	case "nickel":
		return 5
	case "dime":
		return 10
	case "quarter":
		return 20
	}
	return 0
}

func TestCents(t *testing.T) {
	cases := map[string]int{"penny": 1, "nickel": 5, "dime": 10, "quarter": 25, "euro": 0}
	for in, want := range cases {
		if got := cents(in); got != want {
			t.Errorf("cents(%s) = %d, want %d", in, got, want)
		}
	}
}
