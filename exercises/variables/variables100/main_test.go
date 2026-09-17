// variables100
// Make the tests pass!

// I AM NOT DONE
//
// The room numbers are written in the table with leading zeros to look tidy.
// Room "010" is meant to be number 10, but comes out as 8.
// An integer literal with a leading zero is octal.
package main_test

import "testing"

func roomNumbers() []int {
	return []int{007, 010, 012}
}

func TestRoomNumbers(t *testing.T) {
	want := []int{7, 10, 12}
	got := roomNumbers()
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("room %d = %d, want %d", i, got[i], want[i])
		}
	}
}
