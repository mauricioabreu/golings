// functions92
// Make the tests pass!

// I AM NOT DONE
//
// promote must raise the level of an employee by 1.
// The level is unchanged after the call.
// Structs are passed to a function as a copy.
package main_test

import "testing"

type Employee struct {
	Name  string
	Level int
}

func promote(e Employee) {
	e.Level++
}

func promoteTwice(e *Employee) {
	promote(*e)
	promote(*e)
}

func TestPromote(t *testing.T) {
	e := Employee{Name: "Ann", Level: 1}
	promoteTwice(&e)
	if e.Level != 3 {
		t.Errorf("Level = %d, want 3", e.Level)
	}
}
