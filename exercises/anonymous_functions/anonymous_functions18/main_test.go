// anonymous_functions18
// Make the tests pass!

// I AM NOT DONE
//
// withPrefix returns a function adding a prefix to a string.
// Practices capturing a parameter in a closure.
package main_test

import "testing"

func withPrefix(p string) func(string) string {
	return func(s string) string {
		return s + p
	}
}

func TestWithPrefix(t *testing.T) {
	warn := withPrefix("[warn] ")
	if warn("disk") != "[warn] disk" {
		t.Errorf("warn = %q", warn("disk"))
	}
}
