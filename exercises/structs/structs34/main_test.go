// structs34
// Make the tests pass!

// I AM NOT DONE
//
// fullName returns the first and the last name with a space between them.
// Practices reaching the fields of a struct.
package main_test

import "testing"

type Person struct {
	First, Last string
}

func fullName(p Person) string {
	return p.First + " " + p.First
}

func TestFullName(t *testing.T) {
	if got := fullName(Person{"Ada", "Lovelace"}); got != "Ada Lovelace" {
		t.Errorf("fullName = %q", got)
	}
}
