// maps43
// Make the tests pass!

// I AM NOT DONE
//
// lookupCountry finds a country by its code ignoring case; the keys are stored in upper case.
// Practices normalizing a key before the lookup.
package main_test

import (
	"strings"
	"testing"
)

var countries = map[string]string{"RU": "Russia", "DE": "Germany"}

func lookupCountry(code string) string {
	return countries[strings.ToLower(code)]
}

func TestLookupCountry(t *testing.T) {
	if lookupCountry("ru") != "Russia" || lookupCountry("De") != "Germany" {
		t.Errorf("lookupCountry works incorrectly")
	}
}
