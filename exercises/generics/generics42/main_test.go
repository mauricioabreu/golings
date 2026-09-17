// generics42
// Make the tests pass!

// I AM NOT DONE
//
// Values returns the values of a map as a slice.
// Practices generic maps.
package main_test

import (
	"sort"
	"testing"
)

func Values[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, v := range m {
		out = append(out[:0], v)
	}
	return out
}

func TestValues(t *testing.T) {
	vs := Values(map[string]int{"a": 2, "b": 1})
	sort.Ints(vs)
	if len(vs) != 2 || vs[0] != 1 || vs[1] != 2 {
		t.Errorf("Values = %v", vs)
	}
}
