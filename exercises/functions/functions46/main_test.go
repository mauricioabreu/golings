// functions46
// Make the tests pass!

// I AM NOT DONE
//
// apply expects a function, and the result of a call is handed to it instead.
// The code does not compile.
// Practices the difference between f and f().
package main_test

import "testing"

func greet() string { return "hello" }

func apply(f func() string) string {
	return f() + "!"
}

func shout() string {
	return apply(greet())
}

func TestShout(t *testing.T) {
	if got := shout(); got != "hello!" {
		t.Errorf("shout() = %q, want %q", got, "hello!")
	}
}
