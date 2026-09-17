// maps75
// Make the tests pass!

// I AM NOT DONE
//
// The bookings are stored by their starting time. A lookup for the same moment
// written in another time zone finds nothing.
// A time.Time with a different Location is a different map key.
package main_test

import (
	"testing"
	"time"
)

type Bookings map[time.Time]string

func (b Bookings) Book(t time.Time, who string) { b[t] = who }
func (b Bookings) Who(t time.Time) string       { return b[t] }

func TestBookings(t *testing.T) {
	b := Bookings{}
	utc := time.Date(2024, 5, 1, 10, 0, 0, 0, time.UTC)
	b.Book(utc, "ann")
	b.Book(utc.Add(time.Hour), "bob")
	msk := utc.In(time.FixedZone("MSK", 3*3600))
	if got := b.Who(msk); got != "ann" {
		t.Errorf("Who(13:00 MSK) = %q, want ann", got)
	}
	if got := b.Who(utc.Add(time.Hour)); got != "bob" {
		t.Errorf("Who(11:00 UTC) = %q, want bob", got)
	}
}
