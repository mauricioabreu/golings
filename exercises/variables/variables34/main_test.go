// variables34
// Make the tests pass!

// I AM NOT DONE
//
// Celsius is a distinct type based on float64.
// This function must convert degrees Celsius to Fahrenheit.
// Practices named types and conversions between them.
package main_test

import "testing"

type Celsius float64
type Fahrenheit float64

func toFahrenheit(c Celsius) Fahrenheit {
	return c*9/5 + 32
}

func TestToFahrenheit(t *testing.T) {
	if got := toFahrenheit(100); got != 212 {
		t.Errorf("toFahrenheit(100) = %v, want 212", got)
	}
	if got := toFahrenheit(-40); got != -40 {
		t.Errorf("toFahrenheit(-40) = %v, want -40", got)
	}
}
