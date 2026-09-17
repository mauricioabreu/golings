// anonymous_functions52
// Make the tests pass!

// I AM NOT DONE
//
// formatter returns a number formatting literal for a locale:
// "fr" gives "1 234,50" and "en" gives "1,234.50".
// Practices literals configured by the parameters of a factory.
package main_test

import (
	"fmt"
	"strings"
	"testing"
)

func formatter(locale string) func(float64) string {
	group, dec := ",", "."
	if locale == "fr" {
		group, dec = ",", " "
	}
	return func(v float64) string {
		s := fmt.Sprintf("%.2f", v)
		intPart, frac := s[:len(s)-3], s[len(s)-2:]
		var parts []string
		for len(intPart) > 3 {
			parts = append([]string{intPart[len(intPart)-3:]}, parts...)
			intPart = intPart[:len(intPart)-3]
		}
		parts = append([]string{intPart}, parts...)
		return strings.Join(parts, dec) + group + frac
	}
}

func TestFormatter(t *testing.T) {
	if got := formatter("fr")(1234.5); got != "1 234,50" {
		t.Errorf("fr = %q", got)
	}
	if got := formatter("en")(1234567.891); got != "1,234,567.89" {
		t.Errorf("en = %q", got)
	}
}
