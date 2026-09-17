// switch83
// Make the tests pass!

// I AM NOT DONE
//
// weekday returns the name of the day for the number n, where 0 is Monday.
// n may be any integer, negative ones included.
// Practices normalizing a value before a switch.
package main_test

import "testing"

func weekday(n int) string {
	switch n {
	case 0:
		return "Mon"
	case 1:
		return "Tue"
	case 2:
		return "Wed"
	case 3:
		return "Thu"
	case 4:
		return "Fri"
	case 5:
		return "Sat"
	case 7:
		return "Sun"
	}
	return ""
}

func TestWeekday(t *testing.T) {
	cases := map[int]string{0: "Mon", 6: "Sun", 7: "Mon", 15: "Tue", -1: "Sun", -8: "Sun"}
	for in, want := range cases {
		if got := weekday(in); got != want {
			t.Errorf("weekday(%d) = %s, want %s", in, got, want)
		}
	}
}
