// primitive_types67
// Make the tests pass!

// I AM NOT DONE
//
// formatTemp rounds a temperature to a whole number and prints it; the value "-0"
// looks odd and has to be printed as "0".
// float64 has a -0, and it prints with its sign.
package main_test

import (
	"math"
	"strconv"
	"testing"
)

func formatTemp(t float64) string {
	r := math.Round(t)
	return strconv.FormatFloat(r, 'f', 0, 64) + "°"
}

func TestFormatTemp(t *testing.T) {
	cases := map[float64]string{-0.4: "0°", 0.4: "0°", -1.6: "-2°", 21.5: "22°"}
	for in, want := range cases {
		if got := formatTemp(in); got != want {
			t.Errorf("formatTemp(%v) = %q, want %q", in, got, want)
		}
	}
}
