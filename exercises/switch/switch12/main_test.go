// switch12
// Make the tests pass!

// I AM NOT DONE
//
// dayKind takes the day number as an int and compares it against the time.Weekday constants.
// The code does not compile because the types do not match.
// The case values have to be compatible with the type of the tag.
package main_test

import (
	"testing"
	"time"
)

func dayKind(n int) string {
	switch n {
	case time.Saturday, time.Sunday:
		return "weekend"
	}
	return "workday"
}

func TestDayKind(t *testing.T) {
	cases := map[int]string{0: "weekend", 6: "weekend", 1: "workday", 5: "workday"}
	for in, want := range cases {
		if got := dayKind(in); got != want {
			t.Errorf("dayKind(%d) = %s, want %s", in, got, want)
		}
	}
}
