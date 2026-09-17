// if18
// Make the tests pass!

// I AM NOT DONE
//
// validDate checks a day and a month against the length of the month (a non-leap year).
// Practices nested conditions and grouping cases.
package main_test

import "testing"

func validDate(day, month int) bool {
	if month < 1 || month > 12 || day < 1 {
		return false
	}
	maxDay := 31
	if month == 2 {
		maxDay = 29
	}
	return day <= maxDay
}

func TestValidDate(t *testing.T) {
	cases := []struct {
		d, m int
		want bool
	}{{31, 1, true}, {29, 2, false}, {28, 2, true}, {31, 4, false}, {30, 4, true}, {0, 5, false}, {1, 13, false}}
	for _, c := range cases {
		if got := validDate(c.d, c.m); got != c.want {
			t.Errorf("validDate(%d, %d) = %v, want %v", c.d, c.m, got, c.want)
		}
	}
}
