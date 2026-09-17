// maps28
// Make the tests pass!

// I AM NOT DONE
//
// increment passes a counter from a map to a function taking a *int.
// The code does not compile: the address of a map element cannot be taken.
// Map elements may move as the map grows, so they are not addressable.
package main_test

import "testing"

func bump(p *int) { *p += 10 }

func increment(m map[string]int, k string) {
	bump(&m[k])
}

func TestIncrement(t *testing.T) {
	m := map[string]int{"a": 1}
	increment(m, "a")
	increment(m, "b")
	if m["a"] != 11 || m["b"] != 10 {
		t.Errorf("m = %v", m)
	}
}
