// functions20
// Make the tests pass!

// I AM NOT DONE
//
// bindFirst(f, a) must return a one argument function b -> f(a, b).
// Practices partial application.
package main_test

import (
	"strings"
	"testing"
)

func bindFirst(f func(string, string) string, a string) func(string, string) string {
	return f
}

func TestBindFirst(t *testing.T) {
	join := func(a, b string) string { return strings.Join([]string{a, b}, "/") }
	inHome := bindFirst(join, "/home")
	if got := inHome("gopher"); got != "/home/gopher" {
		t.Errorf("inHome(gopher) = %q", got)
	}
}
