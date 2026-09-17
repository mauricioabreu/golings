// functions52
// Make the tests pass!

// I AM NOT DONE
//
// Names must return the list of names so that the caller cannot
// damage the internal state of the registry.
// Returning a slice hands out access to the internal array.
package main_test

import "testing"

type Registry struct{ names []string }

func (r *Registry) Add(n string) { r.names = append(r.names, n) }

func (r *Registry) Names() []string {
	return r.names
}

func TestRegistryNames(t *testing.T) {
	var r Registry
	r.Add("alpha")
	r.Add("beta")
	names := r.Names()
	names[0] = "hacked"
	if got := r.Names()[0]; got != "alpha" {
		t.Errorf("internal state changed: first name = %q, want alpha", got)
	}
}
