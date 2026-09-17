// maps80
// Make the tests pass!

// I AM NOT DONE
//
// defaultHeaders returns the default headers, and the caller adds to them.
// The changes of one call show up in every later one.
// Returning a package level map hands out a shared mutable structure.
package main_test

import "testing"

var defaults = map[string]string{"Accept": "application/json"}

func defaultHeaders() map[string]string {
	return defaults
}

func TestDefaultHeaders(t *testing.T) {
	h1 := defaultHeaders()
	h1["Authorization"] = "secret"
	h2 := defaultHeaders()
	if _, leaked := h2["Authorization"]; leaked || len(h2) != 1 {
		t.Errorf("headers leaked between calls: %v", h2)
	}
}
