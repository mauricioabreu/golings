// if61
// Make the tests pass!

// I AM NOT DONE
//
// parseSwitch reads "on"/"yes"/"1" as true and "off"/"no"/"0" as false,
// ignoring case. Anything else gives def.
// Practices normalizing the input before the comparisons.
package main_test

import (
	"strings"
	"testing"
)

func parseSwitch(s string, def bool) bool {
	if s == "on" || s == "yes" || s == "1" {
		return true
	}
	return false
}

func TestParseSwitch(t *testing.T) {
	_ = strings.ToLower
	cases := []struct {
		in   string
		def  bool
		want bool
	}{{"ON", false, true}, {"yes", false, true}, {"No", true, false}, {"0", true, false}, {"maybe", true, true}, {"", false, false}}
	for _, c := range cases {
		if got := parseSwitch(c.in, c.def); got != c.want {
			t.Errorf("parseSwitch(%q, %v) = %v, want %v", c.in, c.def, got, c.want)
		}
	}
}
