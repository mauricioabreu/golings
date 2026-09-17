// variables96
// Make the tests pass!

// I AM NOT DONE
//
// The package variable greeting must hold "Hello, Gopher".
// The name is set in init(), but greeting is computed earlier.
// Initialization order: package variables first, then init.
package main_test

import "testing"

var name string

func init() {
	name = "Gopher"
}

var greeting = makeGreeting()

func makeGreeting() string {
	return "Hello, " + name
}

func TestGreeting(t *testing.T) {
	if greeting != "Hello, Gopher" {
		t.Errorf("greeting = %q, want %q", greeting, "Hello, Gopher")
	}
}
