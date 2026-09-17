// variables55
// Make the tests pass!

// I AM NOT DONE
//
// makeGreeter must capture the name at the moment the greeting is created.
// Changing the variable name afterwards must not affect the result.
// A closure captures the variable itself, not its value.
package main_test

import "testing"

func makeGreeters() (func() string, func() string) {
	name := "Alice"
	first := func() string { return "Hi, " + name }
	name = "Bob"
	second := func() string { return "Hi, " + name }
	return first, second
}

func TestGreeters(t *testing.T) {
	first, second := makeGreeters()
	if got := first(); got != "Hi, Alice" {
		t.Errorf("first() = %q, want %q", got, "Hi, Alice")
	}
	if got := second(); got != "Hi, Bob" {
		t.Errorf("second() = %q, want %q", got, "Hi, Bob")
	}
}
