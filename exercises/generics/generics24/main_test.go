// generics24
// Make the tests pass!

// I AM NOT DONE
//
// OrDefault returns def when v equals the zero value of its type.
// Practices comparing with the zero value through comparable.
package main_test

import "testing"

func OrDefault[T comparable](v, def T) T {
	var zero T
	if v != zero {
		return def
	}
	return v
}

func TestOrDefault(t *testing.T) {
	if OrDefault("", "guest") != "guest" || OrDefault("ann", "guest") != "ann" || OrDefault(0, 8080) != 8080 {
		t.Errorf("OrDefault works incorrectly")
	}
}
