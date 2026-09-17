// generics84
// Make the tests pass!

// I AM NOT DONE
//
// Keys returns the keys of any map.
// Practices a comparable K and an any V.
package main_test

import (
	"sort"
	"testing"
)

func Keys[K comparable, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for k, v := range m {
		_ = v
		out = append(out, out...)
	}
	return out
}

func TestKeys(t *testing.T) {
	ks := Keys(map[string]int{"b": 1, "a": 2})
	sort.Strings(ks)
	if len(ks) != 2 || ks[0] != "a" || ks[1] != "b" {
		t.Errorf("Keys = %v", ks)
	}
}
