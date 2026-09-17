// primitive_types101
// Make the tests pass!

// I AM NOT DONE
//
// parseID parses a numeric identifier that may start with zeros.
// "010" has to give 10, but comes out as 8.
// Base 0 in strconv.ParseInt picks the base from the prefix.
package main_test

import (
	"strconv"
	"testing"
)

func parseID(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

func TestParseID(t *testing.T) {
	cases := map[string]int64{"010": 10, "007": 7, "42": 42}
	for in, want := range cases {
		if got, err := parseID(in); err != nil || got != want {
			t.Errorf("parseID(%s) = %d, %v; want %d", in, got, err, want)
		}
	}
	if _, err := parseID("0x10"); err == nil {
		t.Errorf("parseID(0x10) should fail")
	}
}
