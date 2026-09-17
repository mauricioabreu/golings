// functions39
// Make the tests pass!

// I AM NOT DONE
//
// GreeterFunc is an adapter that lets a function be used as a Greeter.
// The type has no method yet, so the function does not satisfy the interface.
// Practices methods on function types, the way http.HandlerFunc works.
package main_test

import "testing"

type Greeter interface {
	Greet(name string) string
}

type GreeterFunc func(string) string

func welcome(g Greeter) string {
	return g.Greet("gopher")
}

func polite(name string) string { return "Good day, " + name }

func run() string {
	return welcome(polite)
}

func TestRun(t *testing.T) {
	if got := run(); got != "Good day, gopher" {
		t.Errorf("run() = %q", got)
	}
}
