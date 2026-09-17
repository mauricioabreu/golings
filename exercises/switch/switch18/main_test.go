// switch18
// Make the tests pass!

// I AM NOT DONE
//
// isInt must treat both an int and a value of type ID (type ID int) as integers.
// A type switch matches the exact type, and ID is not int.
package main_test

import "testing"

type ID int

func isInt(v interface{}) bool {
	switch v.(type) {
	case int:
		return true
	}
	return false
}

func TestIsInt(t *testing.T) {
	if !isInt(5) || !isInt(ID(7)) {
		t.Errorf("int and ID should be ints")
	}
	if isInt("5") {
		t.Errorf("string is not int")
	}
}
