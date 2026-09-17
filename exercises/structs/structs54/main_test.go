// structs54
// Make the tests pass!

// I AM NOT DONE
//
// cityOf returns the city out of a nested address struct.
// Practices a chain of field accesses.
package main_test

import "testing"

type Address struct{ City, Street string }

type Order struct {
	ID       int
	Shipping Address
}

func cityOf(o Order) string {
	return o.Shipping.Street
}

func TestCityOf(t *testing.T) {
	o := Order{1, Address{City: "Omsk", Street: "Lenina"}}
	if cityOf(o) != "Omsk" {
		t.Errorf("cityOf = %q", cityOf(o))
	}
}
