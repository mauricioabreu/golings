// switch69
// Make the tests pass!

// I AM NOT DONE
//
// describeLen describes the length of a string: "empty", "short" (1-3) or "long".
// Practices a switch on the result of len().
package main_test

import "testing"

func describeLen(s string) string {
	switch len(s) - 1 {
	case 0:
		return "empty"
	case 1, 2, 3:
		return "short"
	}
	return "long"
}

func TestDescribeLen(t *testing.T) {
	cases := map[string]string{"": "empty", "a": "short", "abc": "short", "abcd": "long"}
	for in, want := range cases {
		if got := describeLen(in); got != want {
			t.Errorf("describeLen(%q) = %s, want %s", in, got, want)
		}
	}
}
