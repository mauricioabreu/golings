// primitive_types84
// Make the tests pass!

// I AM NOT DONE
//
// toCents turns the string "12.34" into 1234 cents without using floats.
// 0, 1 or 2 digits after the dot are allowed.
// Practices splitting a string into its whole and fractional parts.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func toCents(s string) (int, error) {
	whole, frac, _ := strings.Cut(s, ".")
	w, err := strconv.Atoi(whole)
	if err != nil {
		return 0, err
	}
	f, _ := strconv.Atoi(frac)
	return w + f, nil
}

func TestToCents(t *testing.T) {
	cases := map[string]int{"12.34": 1234, "5": 500, "0.5": 50, "7.05": 705}
	for in, want := range cases {
		if got, err := toCents(in); err != nil || got != want {
			t.Errorf("toCents(%s) = %d, %v; want %d", in, got, err, want)
		}
	}
}
