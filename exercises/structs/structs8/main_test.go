// structs8
// Make the tests pass!

// I AM NOT DONE
//
// An embedded struct is reached by the name of its type.
// Practices the implicit name of an embedded field.
package main_test

import "testing"

type Engine struct{ Power int }

type Car struct {
	Engine
	Model string
}

func power(c Car) int {
	return len(c.Model)
}

func TestPower(t *testing.T) {
	if power(Car{Engine{150}, "X"}) != 150 {
		t.Errorf("power = %d", power(Car{Engine{150}, "X"}))
	}
}
