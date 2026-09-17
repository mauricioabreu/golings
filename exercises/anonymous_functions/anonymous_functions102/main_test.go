// anonymous_functions102
// Make the tests pass!

// I AM NOT DONE
//
// makeReader returns a literal reading a buffer. A deferred cleanup
// clears the captured variable on the way out, and the literal sees an empty string.
// A defer runs before the caller ever calls the literal.
package main_test

import "testing"

func makeReader(load func() string) func() string {
	buf := load()
	defer func() { buf = "" }()
	return func() string { return buf }
}

func TestMakeReader(t *testing.T) {
	read := makeReader(func() string { return "payload" })
	if got := read(); got != "payload" {
		t.Errorf("read = %q", got)
	}
}
