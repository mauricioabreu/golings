// arrays8
// Make the tests pass!

// I AM NOT DONE
//
// The message table is built with an indexed literal.
// The code does not compile: the same index is given twice.
// The indexes in an array literal have to be distinct.
package main_test

import "testing"

const (
	OK       = 0
	NotFound = 1
	Denied   = 2
)

var text = [3]string{
	OK:     "ok",
	OK:     "not found",
	Denied: "denied",
}

func TestText(t *testing.T) {
	if text[OK] != "ok" || text[NotFound] != "not found" || text[Denied] != "denied" {
		t.Errorf("text = %v", text)
	}
}
