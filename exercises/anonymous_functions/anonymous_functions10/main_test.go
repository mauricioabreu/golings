// anonymous_functions10
// Make the tests pass!

// I AM NOT DONE
//
// greeter returns a literal greeting the captured name.
// Practices capturing the parameter of a factory.
package main_test

import "testing"

func greeter(name string) func() string {
	return func() string {
		return "Hello, world"
	}
}

func TestGreeter(t *testing.T) {
	if greeter("Go")() != "Hello, Go" {
		t.Errorf("greeter = %q", greeter("Go")())
	}
}
