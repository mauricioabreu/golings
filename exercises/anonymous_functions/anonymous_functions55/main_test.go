// anonymous_functions55
// Make the tests pass!

// I AM NOT DONE
//
// withDefaults returns a function filling a request in with default values.
// The literal changes the shared defaults object, and later requests get somebody else's values.
// Capturing a pointer gives access to a shared mutable object.
package main_test

import "testing"

type Req struct {
	Timeout int
	Path    string
}

func withDefaults(def *Req) func(path string) Req {
	return func(path string) Req {
		r := def
		r.Path = path
		return *r
	}
}

func TestWithDefaults(t *testing.T) {
	def := &Req{Timeout: 30}
	mk := withDefaults(def)
	a := mk("/a")
	b := mk("/b")
	if a.Path != "/a" || b.Path != "/b" || def.Path != "" {
		t.Errorf("a=%+v b=%+v def=%+v", a, b, *def)
	}
}
