// variables22
// Make the tests pass!

// I AM NOT DONE
//
// The days of the week are declared with iota, and the String method must return their names.
// The numbering and the name table do not line up.
// Practices iota, named types and methods on them.
package main_test

import "testing"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
)

var weekdayNames = [...]string{"Mon", "Tue", "Wed"}

func (d Weekday) String() string {
	return weekdayNames[d]
}

func TestWeekdayString(t *testing.T) {
	cases := map[Weekday]string{Monday: "Mon", Tuesday: "Tue", Wednesday: "Wed"}
	for d, want := range cases {
		if got := d.String(); got != want {
			t.Errorf("Weekday(%d).String() = %s, want %s", int(d), got, want)
		}
	}
}
