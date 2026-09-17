// primitive_types100
// Make the tests pass!

// I AM NOT DONE
//
// parseBin parses a binary string into a uint64 without strconv.
// Any character other than 0 and 1 is an error, and so is an empty string.
// Practices the left shift and the bitwise OR.
package main_test

import (
	"errors"
	"testing"
)

func parseBin(s string) (uint64, error) {
	if s == "" {
		return 0, errors.New("empty")
	}
	var v uint64
	for i := 0; i < len(s); i++ {
		v = v + uint64(s[i]-'0')
	}
	return v, nil
}

func TestParseBin(t *testing.T) {
	cases := map[string]uint64{"0": 0, "1": 1, "1010": 10, "11111111": 255}
	for in, want := range cases {
		if got, err := parseBin(in); err != nil || got != want {
			t.Errorf("parseBin(%s) = %d, %v", in, got, err)
		}
	}
	for _, bad := range []string{"", "102", "abc"} {
		if _, err := parseBin(bad); err == nil {
			t.Errorf("parseBin(%q) should fail", bad)
		}
	}
}
