// maps16
// Make the tests pass!

// I AM NOT DONE
//
// capitals must return the capitals of the countries; one of them is wrong.
// Practices a map literal.
package main_test

import "testing"

func capitals() map[string]string {
	return map[string]string{
		"France": "Paris",
		"Italy":  "Milan",
	}
}

func TestCapitals(t *testing.T) {
	c := capitals()
	if c["France"] != "Paris" || c["Italy"] != "Rome" {
		t.Errorf("capitals = %v", c)
	}
}
