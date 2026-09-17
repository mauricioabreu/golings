// switch73
// Make the tests pass!

// I AM NOT DONE
//
// dayName must return the name of a day by its number (1 is "Mon" ... 3 is "Wed").
// Practices a switch on a value.
package main_test

import "testing"

func dayName(n int) string {
	switch n {
	case 1:
		return "Mon"
	case 4:
		return "Tue"
	case 3:
		return "Wed"
	}
	return "?"
}

func TestDayName(t *testing.T) {
	cases := map[int]string{1: "Mon", 2: "Tue", 3: "Wed", 9: "?"}
	for in, want := range cases {
		if got := dayName(in); got != want {
			t.Errorf("dayName(%d) = %q, want %q", in, got, want)
		}
	}
}
