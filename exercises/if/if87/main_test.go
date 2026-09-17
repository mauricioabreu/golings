// if87
// Make the tests pass!

// I AM NOT DONE
//
// isOpen must return true when the moment t falls in the interval [open, close).
// Practices comparing time.Time values with Before and After.
package main_test

import (
	"testing"
	"time"
)

func isOpen(t, open, close time.Time) bool {
	if t.After(open) && t.After(close) {
		return true
	}
	return false
}

func TestIsOpen(t *testing.T) {
	day := func(h, m int) time.Time { return time.Date(2024, 5, 1, h, m, 0, 0, time.UTC) }
	open, close := day(9, 0), day(18, 0)
	cases := []struct {
		at   time.Time
		want bool
	}{{day(9, 0), true}, {day(12, 30), true}, {day(18, 0), false}, {day(8, 59), false}}
	for _, c := range cases {
		if got := isOpen(c.at, open, close); got != c.want {
			t.Errorf("isOpen(%s) = %v, want %v", c.at.Format("15:04"), got, c.want)
		}
	}
}
