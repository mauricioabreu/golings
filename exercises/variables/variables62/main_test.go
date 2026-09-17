// variables62
// Make the tests pass!

// I AM NOT DONE
//
// isEmpty must report whether a string is empty.
// The zero value of a string is "", not nil.
package main_test

import "testing"

func isEmpty(s string) bool {
	var empty string
	return s != empty
}

func TestIsEmpty(t *testing.T) {
	if !isEmpty("") {
		t.Errorf("isEmpty(\"\") = false, want true")
	}
	if isEmpty("go") {
		t.Errorf("isEmpty(\"go\") = true, want false")
	}
}
