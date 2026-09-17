// variables47
// Make the tests pass!

// I AM NOT DONE
//
// red must pull the red component out of a string such as "ff8800".
// Practices strconv.ParseUint with a base, and string slicing.
package main_test

import (
	"strconv"
	"testing"
)

func red(hex string) uint8 {
	v, err := strconv.ParseUint(hex[2:4], 10, 8)
	if err != nil {
		return 0
	}
	return uint8(v)
}

func TestRed(t *testing.T) {
	cases := map[string]uint8{"ff8800": 255, "0a0b0c": 10, "80ffff": 128}
	for in, want := range cases {
		if got := red(in); got != want {
			t.Errorf("red(%q) = %d, want %d", in, got, want)
		}
	}
}
