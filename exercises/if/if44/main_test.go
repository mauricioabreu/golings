// if44
// Make the tests pass!

// I AM NOT DONE
//
// firstValid parses two numbers and returns the FIRST one when both are valid.
// Right now it returns the second one.
// Variables from an else if initializer shadow those from the if.
package main_test

import (
	"strconv"
	"testing"
)

func firstValid(a, b string) int {
	if n, err := strconv.Atoi(a); err != nil {
		return -1
	} else if n, err := strconv.Atoi(b); err != nil {
		return -2
	} else {
		return n
	}
}

func TestFirstValid(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{{"1", "2", 1}, {"7", "x", -2}, {"x", "2", -1}}
	for _, c := range cases {
		if got := firstValid(c.a, c.b); got != c.want {
			t.Errorf("firstValid(%s, %s) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
