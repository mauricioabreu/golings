// anonymous_functions51
// Make the tests pass!

// I AM NOT DONE
//
// handler is built capturing the db variable, and the variable is then set to nil
// on "close". The handler has to use the connection it was built with.
// A literal captures the variable, not the current value of the pointer.
package main_test

import "testing"

type DB struct{ Name string }

func setup() func() string {
	db := &DB{Name: "main"}
	handler := func() string { return db.Name }
	db = nil
	return handler
}

func TestSetup(t *testing.T) {
	if got := setup()(); got != "main" {
		t.Errorf("handler = %q", got)
	}
}
