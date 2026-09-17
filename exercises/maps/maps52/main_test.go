// maps52
// Make the tests pass!

// I AM NOT DONE
//
// seats keeps track of who sits in a (row, seat) place.
// Practices a struct as a map key.
package main_test

import "testing"

type Seat struct{ Row, Col int }

func whoSits(m map[Seat]string, row, col int) string {
	return m[Seat{col, row}]
}

func TestWhoSits(t *testing.T) {
	m := map[Seat]string{{1, 2}: "ann", {2, 1}: "bob"}
	if whoSits(m, 1, 2) != "ann" || whoSits(m, 2, 1) != "bob" {
		t.Errorf("whoSits works incorrectly")
	}
}
