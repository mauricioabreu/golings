// arrays5
// Make the tests pass!

// I AM NOT DONE
//
// daysIn returns the number of days in a month (1-12) of a non-leap year.
// Practices an array as a lookup table.
package main_test

import "testing"

var monthDays = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func daysIn(month int) int {
	return monthDays[month]
}

func TestDaysIn(t *testing.T) {
	cases := map[int]int{1: 31, 2: 28, 4: 30, 12: 31}
	for in, want := range cases {
		if got := daysIn(in); got != want {
			t.Errorf("daysIn(%d) = %d, want %d", in, got, want)
		}
	}
}
