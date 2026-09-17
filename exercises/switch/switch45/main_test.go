// switch45
// Make the tests pass!

// I AM NOT DONE
//
// isWeekend must return true for "sat" and "sun".
// Practices a list of values in a case.
package main_test

import "testing"

func isWeekend(day string) bool {
	switch day {
	case "sat", "fri":
		return true
	default:
		return false
	}
}

func TestIsWeekend(t *testing.T) {
	cases := map[string]bool{"sat": true, "sun": true, "fri": false, "mon": false}
	for in, want := range cases {
		if got := isWeekend(in); got != want {
			t.Errorf("isWeekend(%s) = %v, want %v", in, got, want)
		}
	}
}
