// maps98
// Make the tests pass!

// I AM NOT DONE
//
// The registry keeps handlers in a map[interface{}]string. Registration uses an int
// while the lookup uses an int64 from an outside source, and the handler is not found.
// interface{} keys are equal only when both the type and the value match.
package main_test

import "testing"

type Registry map[interface{}]string

func (r Registry) Register(code int, name string) {
	r[code] = name
}

func (r Registry) Lookup(code int64) string {
	return r[code]
}

func TestRegistry(t *testing.T) {
	r := Registry{}
	r.Register(200, "ok")
	r.Register(404, "not found")
	var fromWire int64 = 404
	if got := r.Lookup(fromWire); got != "not found" {
		t.Errorf("Lookup(404) = %q, want not found", got)
	}
}
