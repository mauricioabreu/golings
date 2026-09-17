// if30
// Make the tests pass!

// I AM NOT DONE
//
// isWeekend must return true for "sat" and "sun".
// Practices the logical operators && and ||.
package main_test

import "testing"

func isWeekend(day string) bool {
	if day == "sat" && day == "sun" {
		return true
	}
	return false
}

func TestIsWeekend(t *testing.T) {
	cases := map[string]bool{"sat": true, "sun": true, "mon": false, "fri": false}
	for in, want := range cases {
		if got := isWeekend(in); got != want {
			t.Errorf("isWeekend(%q) = %v, want %v", in, got, want)
		}
	}
}
