// functions65
// Make the tests pass!

// I AM NOT DONE
//
// greet(name, punct...) uses the first optional argument as the punctuation
// and falls back to ".".
// A variadic parameter as a way to offer an optional argument.
package main_test

import "testing"

func greet(name string, punct ...string) string {
	p := "."
	if len(punct) > 1 {
		p = punct[0]
	}
	return "Hi, " + name + p
}

func TestGreet(t *testing.T) {
	if got := greet("Ann"); got != "Hi, Ann." {
		t.Errorf("greet(Ann) = %q", got)
	}
	if got := greet("Bob", "!"); got != "Hi, Bob!" {
		t.Errorf("greet(Bob, !) = %q", got)
	}
}
