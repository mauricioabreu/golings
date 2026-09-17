// anonymous_functions43
// Make the tests pass!

// I AM NOT DONE
//
// strings.Map asks for a func(rune) rune. The literal is declared with another type.
// The signature of a literal has to match the expected one exactly.
package main_test

import (
	"strings"
	"testing"
)

func upperA(s string) string {
	return strings.Map(func(r rune) byte {
		if r == 'a' {
			return 'A'
		}
		return r
	}, s)
}

func TestUpperA(t *testing.T) {
	if upperA("banana") != "bAnAnA" {
		t.Errorf("upperA = %q", upperA("banana"))
	}
}
