// anonymous_functions63
// Make the tests pass!

// I AM NOT DONE
//
// The handlers are kept as interface{}. The literal has the unnamed type
// func(string) string, so an assertion to Handler never succeeds.
// The dynamic type of a literal is an unnamed function type.
package main_test

import "testing"

type Handler func(string) string

func call(h interface{}, arg string) string {
	if f, ok := h.(Handler); ok {
		return f(arg)
	}
	return "not a handler"
}

func TestCall(t *testing.T) {
	var h interface{} = func(s string) string { return "hi " + s }
	if got := call(h, "go"); got != "hi go" {
		t.Errorf("call = %q", got)
	}
	if got := call(Handler(func(s string) string { return s }), "x"); got != "x" {
		t.Errorf("call(Handler) = %q", got)
	}
}
