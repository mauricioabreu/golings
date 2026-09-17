// maps13
// Make the tests pass!

// I AM NOT DONE
//
// invert swaps the keys and the values, the values being distinct.
// Practices writing into a new map in a loop.
package main_test

import (
	"reflect"
	"testing"
)

func invert(m map[string]int) map[int]string {
	out := make(map[int]string, len(m))
	for k, v := range m {
		out[len(k)] = k
	}
	return out
}

func TestInvert(t *testing.T) {
	got := invert(map[string]int{"one": 1, "two": 2})
	if !reflect.DeepEqual(got, map[int]string{1: "one", 2: "two"}) {
		t.Errorf("invert = %v", got)
	}
}
