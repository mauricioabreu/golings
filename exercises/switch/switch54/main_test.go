// switch54
// Make the tests pass!

// I AM NOT DONE
//
// parseDur turns strings such as "30s", "5m" and "2h" into seconds.
// Practices a switch on the last character of a string.
package main_test

import (
	"errors"
	"strconv"
	"testing"
)

func parseDur(s string) (int, error) {
	if len(s) < 2 {
		return 0, errors.New("too short")
	}
	n, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return 0, err
	}
	switch s[len(s)-1] {
	case 's':
		return n, nil
	case 'm':
		return n * 3600, nil
	}
	return n, nil
}

func TestParseDur(t *testing.T) {
	cases := map[string]int{"30s": 30, "5m": 300, "2h": 7200}
	for in, want := range cases {
		if got, err := parseDur(in); err != nil || got != want {
			t.Errorf("parseDur(%s) = %d, %v; want %d", in, got, err, want)
		}
	}
	for _, bad := range []string{"5d", "s", "xm"} {
		if _, err := parseDur(bad); err == nil {
			t.Errorf("parseDur(%s) should fail", bad)
		}
	}
}
