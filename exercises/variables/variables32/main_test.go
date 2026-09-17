// variables32
// Make the tests pass!

// I AM NOT DONE
//
// nextID must hand out consecutive numbers: 1, 2, 3 and so on.
// Practices package level variables and the order of the increment.
package main_test

import "testing"

var lastID int

func nextID() int {
	id := lastID
	lastID++
	return id
}

func TestNextID(t *testing.T) {
	lastID = 0
	for want := 1; want <= 3; want++ {
		if got := nextID(); got != want {
			t.Errorf("nextID() = %d, want %d", got, want)
		}
	}
}
