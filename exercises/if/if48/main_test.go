// if48
// Make the tests pass!

// I AM NOT DONE
//
// newer must compare build numbers written as strings with no leading zeros.
// "10" has to be newer than "9".
// Strings compare lexicographically, byte by byte.
package main_test

import (
	"strconv"
	"testing"
)

func newer(a, b string) bool {
	if a > b {
		return true
	}
	return false
}

func TestNewer(t *testing.T) {
	_ = strconv.Atoi
	cases := []struct {
		a, b string
		want bool
	}{{"10", "9", true}, {"9", "10", false}, {"100", "99", true}, {"2", "1", true}, {"5", "5", false}}
	for _, c := range cases {
		if got := newer(c.a, c.b); got != c.want {
			t.Errorf("newer(%s, %s) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
