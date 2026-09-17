// variables54
// Make the tests pass!

// I AM NOT DONE
//
// parseAmount must return a number along with whether it was parsed successfully.
// Practices strconv.ParseFloat and working with several results.
package main_test

import (
	"strconv"
	"testing"
)

func parseAmount(s string) (float64, bool) {
	v, err := strconv.ParseFloat(s, 32)
	return v, err != nil
}

func TestParseAmount(t *testing.T) {
	if v, ok := parseAmount("19.99"); !ok || v != 19.99 {
		t.Errorf("parseAmount(19.99) = %v, %v", v, ok)
	}
	if _, ok := parseAmount("12,5"); ok {
		t.Errorf("parseAmount(12,5) should fail")
	}
}
