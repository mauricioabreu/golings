// maps103
// Make the tests pass!

// I AM NOT DONE
//
// uniqueInOrder returns the distinct strings in order of first appearance.
// The result is built from the keys of a map, and the order is lost.
// A map does not keep the insertion order.
package main_test

import (
	"reflect"
	"testing"
)

func uniqueInOrder(items []string) []string {
	seen := map[string]bool{}
	for _, it := range items {
		seen[it] = true
	}
	var out []string
	for it := range seen {
		out = append(out, it)
	}
	return out
}

func TestUniqueInOrder(t *testing.T) {
	in := []string{"z", "a", "z", "m", "a", "q", "b", "c"}
	for i := 0; i < 30; i++ {
		if got := uniqueInOrder(in); !reflect.DeepEqual(got, []string{"z", "a", "m", "q", "b", "c"}) {
			t.Fatalf("uniqueInOrder = %v", got)
		}
	}
}
