// if63
// Make the tests pass!

// I AM NOT DONE
//
// parseOr must return the parsed number, or def.
// The code does not compile: a variable from the if is used after the block.
// Variables from an if initializer are only visible inside the if and its else.
package main_test

import (
	"strconv"
	"testing"
)

func parseOr(s string, def int) int {
	if n, err := strconv.Atoi(s); err != nil {
		return def
	}
	return n
}

func TestParseOr(t *testing.T) {
	if got := parseOr("12", 0); got != 12 {
		t.Errorf("parseOr(12) = %d", got)
	}
	if got := parseOr("x", 5); got != 5 {
		t.Errorf("parseOr(x, 5) = %d", got)
	}
}
