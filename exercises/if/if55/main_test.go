// if55
// Make the tests pass!

// I AM NOT DONE
//
// isInteger must return true for a value that arrived as a plain integer constant.
// Right now the check for an int64 never matches.
// An untyped constant put in an interface{} takes the type int.
package main_test

import "testing"

func isInteger(v interface{}) bool {
	if _, ok := v.(int64); ok {
		return true
	}
	return false
}

func TestIsInteger(t *testing.T) {
	if !isInteger(42) {
		t.Errorf("isInteger(42) = false, want true")
	}
	if isInteger("42") || isInteger(4.2) {
		t.Errorf("strings and floats are not integers")
	}
}
