// variables28
// Make the tests pass!

// I AM NOT DONE
//
// The constant SecondsPerDay must hold the number of seconds in a day.
// Practices constant expressions evaluated at compile time.
package main_test

import "testing"

const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
)

const SecondsPerDay = SecondsPerMinute * MinutesPerHour

func TestSecondsPerDay(t *testing.T) {
	if SecondsPerDay != 86400 {
		t.Errorf("SecondsPerDay = %d, want 86400", SecondsPerDay)
	}
}
