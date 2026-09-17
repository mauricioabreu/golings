// primitive_types57
// Make the tests pass!

// I AM NOT DONE
//
// censor must replace every digit in a string with '*'.
// The bytes are changed, yet the function returns the original string.
// []byte(s) makes a copy, and the string itself does not change.
package main_test

import "testing"

func censor(s string) string {
	b := []byte(s)
	for i := range b {
		if b[i] >= '0' && b[i] <= '9' {
			b[i] = '*'
		}
	}
	return s
}

func TestCensor(t *testing.T) {
	if got := censor("card 1234"); got != "card ****" {
		t.Errorf("censor = %q, want %q", got, "card ****")
	}
}
