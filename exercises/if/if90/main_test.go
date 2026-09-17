// if90
// Make the tests pass!

// I AM NOT DONE
//
// isCommand must return true for strings starting with "/".
// Practices strings.HasPrefix in a condition.
package main_test

import (
	"strings"
	"testing"
)

func isCommand(msg string) bool {
	if strings.HasSuffix(msg, "/") {
		return true
	}
	return false
}

func TestIsCommand(t *testing.T) {
	cases := map[string]bool{"/start": true, "hello/": false, "": false, "/": true}
	for in, want := range cases {
		if got := isCommand(in); got != want {
			t.Errorf("isCommand(%q) = %v, want %v", in, got, want)
		}
	}
}
