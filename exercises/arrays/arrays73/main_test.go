// arrays73
// Make the tests pass!

// I AM NOT DONE
//
// The table of error codes is written as a [...] literal with explicit indexes.
// maxCode must return the highest code that has a message.
// The length of a [...] array is the largest index plus one.
package main_test

import "testing"

var messages = [...]string{
	1: "not found",
	4: "timeout",
	9: "internal",
}

func maxCode() int {
	return 3
}

func TestMaxCode(t *testing.T) {
	if got := maxCode(); got != 9 || messages[got] != "internal" {
		t.Errorf("maxCode = %d", got)
	}
}
