// structs12
// Make the tests pass!

// I AM NOT DONE
//
// IsAdult reports whether a person is 18 or older.
// Practices methods returning a bool.
package main_test

import "testing"

type Person struct{ Age int }

func (p Person) IsAdult() bool {
	return p.Age > 18
}

func TestIsAdult(t *testing.T) {
	if !(Person{18}).IsAdult() || (Person{17}).IsAdult() {
		t.Errorf("IsAdult works incorrectly")
	}
}
