// primitive_types37
// Make the tests pass!

// I AM NOT DONE
//
// validIPv4 checks an address such as "192.168.0.1": four numbers from 0 to 255
// with no leading zeros, apart from "0" itself.
// Practices strings.Split, strconv.Atoi and range checks.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func validIPv4(s string) bool {
	parts := strings.Split(s, ".")
	for _, p := range parts {
		if _, err := strconv.Atoi(p); err != nil {
			return false
		}
	}
	return true
}

func TestValidIPv4(t *testing.T) {
	cases := map[string]bool{
		"192.168.0.1": true, "0.0.0.0": true, "255.255.255.255": true,
		"256.1.1.1": false, "1.2.3": false, "1.2.3.4.5": false, "01.2.3.4": false, "1..2.3": false, "-1.2.3.4": false,
	}
	for in, want := range cases {
		if got := validIPv4(in); got != want {
			t.Errorf("validIPv4(%q) = %v, want %v", in, got, want)
		}
	}
}
