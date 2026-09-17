// switch50
// Make the tests pass!

// I AM NOT DONE
//
// isWorkday uses time.Weekday. Saturday and Sunday are the weekend.
// Practices a switch over the values of a named type from the standard library.
package main_test

import (
	"testing"
	"time"
)

func isWorkday(d time.Weekday) bool {
	switch d {
	case time.Saturday, time.Monday:
		return false
	}
	return true
}

func TestIsWorkday(t *testing.T) {
	cases := map[time.Weekday]bool{time.Monday: true, time.Friday: true, time.Saturday: false, time.Sunday: false}
	for in, want := range cases {
		if got := isWorkday(in); got != want {
			t.Errorf("isWorkday(%s) = %v, want %v", in, got, want)
		}
	}
}
