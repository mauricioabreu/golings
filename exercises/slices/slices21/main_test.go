// slices21
// Make the tests pass!

// I AM NOT DONE
//
// difference returns the elements of a that are not in b.
// Practices a set and filtering a slice.
package main_test

import (
	"reflect"
	"testing"
)

func difference(a, b []string) []string {
	set := map[string]bool{}
	for _, v := range b {
		set[v] = true
	}
	var out []string
	for _, v := range a {
		if set[v] {
			out = append(out, v)
		}
	}
	return a
}

func TestDifference(t *testing.T) {
	if got := difference([]string{"a", "b", "c", "d"}, []string{"b", "d", "x"}); !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("difference = %v", got)
	}
}
