// functions53
// Make the tests pass!

// I AM NOT DONE
//
// chain(h, m1, m2) must return a handler with m1 as the outermost layer.
// The call must produce "m1(m2(handler))".
// Practices composing wrapper functions.
package main_test

import "testing"

type Handler func() string
type Middleware func(Handler) Handler

func named(name string) Middleware {
	return func(next Handler) Handler {
		return func() string { return name + "(" + next() + ")" }
	}
}

func chain(h Handler, mws ...Middleware) Handler {
	for _, m := range mws {
		h = m(h)
	}
	return h
}

func TestChain(t *testing.T) {
	h := chain(func() string { return "handler" }, named("m1"), named("m2"))
	if got := h(); got != "m1(m2(handler))" {
		t.Errorf("chain() = %q, want %q", got, "m1(m2(handler))")
	}
}
