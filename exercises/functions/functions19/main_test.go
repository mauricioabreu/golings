// functions19
// Make the tests pass!

// I AM NOT DONE
//
// The handler has to match the Handler type, but it has no use for id.
// The code does not compile: the signature does not match the type.
// Practices function types and the blank identifier in a parameter list.
package main_test

import "testing"

type Handler func(id int, name string) string

func run(h Handler) string {
	return h(42, "gopher")
}

func hello(name string) string {
	return "hello " + name
}

func TestRunHello(t *testing.T) {
	if got := run(hello); got != "hello gopher" {
		t.Errorf("run(hello) = %q, want %q", got, "hello gopher")
	}
}
