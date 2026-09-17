// slices39
// Make the tests pass!

// I AM NOT DONE
//
// unique removes the repeats, keeping the order of the first occurrences.
// Practices a slice together with a map used as a set.
package main_test

import (
	"reflect"
	"testing"
)

func unique(s []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, v := range s {
		if seen[v] {
			out = append(out, v)
		}
	}
	return out
}

func TestUnique(t *testing.T) {
	if got := unique([]string{"b", "a", "b", "c", "a"}); !reflect.DeepEqual(got, []string{"b", "a", "c"}) {
		t.Errorf("unique = %v", got)
	}
}
