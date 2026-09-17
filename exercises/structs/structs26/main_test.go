// structs26
// Make the tests pass!

// I AM NOT DONE
//
// The embedded Address struct gives Company a City field directly.
// Practices embedding and field promotion.
package main_test

import "testing"

type Address struct{ City string }

type Company struct {
	Name string
	Addr Address
}

func TestCompanyCity(t *testing.T) {
	c := Company{Name: "Acme"}
	c.Address.City = "Kazan"
	if c.City != "Kazan" {
		t.Errorf("City = %q", c.City)
	}
}
