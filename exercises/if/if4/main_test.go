// if4
// Make the tests pass!

// I AM NOT DONE
//
// sameMoment must treat moments written in different time zones as equal.
// A time.Time cannot be compared reliably with ==.
package main_test

import (
	"testing"
	"time"
)

func sameMoment(a, b time.Time) bool {
	if a == b {
		return true
	}
	return false
}

func TestSameMoment(t *testing.T) {
	utc := time.Date(2024, 3, 1, 12, 0, 0, 0, time.UTC)
	msk := utc.In(time.FixedZone("MSK", 3*3600))
	if !sameMoment(utc, msk) {
		t.Errorf("12:00 UTC and 15:00 MSK are the same moment")
	}
	if sameMoment(utc, utc.Add(time.Second)) {
		t.Errorf("different moments should differ")
	}
}
