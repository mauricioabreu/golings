// arrays100
// Make the tests pass!

// I AM NOT DONE
//
// weekdayName returns the name of a day for a number coming from user input,
// or an error. For an invalid number the program panics.
// Going out of the bounds of an array with a variable index is a run time panic.
package main_test

import (
	"errors"
	"testing"
)

var days = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func weekdayName(n int) (string, error) {
	if n > len(days) {
		return "", errors.New("bad day")
	}
	return days[n-1], nil
}

func TestWeekdayName(t *testing.T) {
	if d, err := weekdayName(1); err != nil || d != "Mon" {
		t.Errorf("weekdayName(1) = %s, %v", d, err)
	}
	for _, bad := range []int{0, -3, 8} {
		if _, err := weekdayName(bad); err == nil {
			t.Errorf("weekdayName(%d) should fail", bad)
		}
	}
}
