// range25
// Make the tests pass!

// I AM NOT DONE
//
// listKeys returns the keys of a map sorted alphabetically.
// Practices a range over a map followed by a sort.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func listKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	_ = sort.Strings
	return keys
}

func TestListKeys(t *testing.T) {
	if got := listKeys(map[string]int{"z": 1, "a": 2, "m": 3}); !reflect.DeepEqual(got, []string{"a", "m", "z"}) {
		t.Errorf("listKeys = %v", got)
	}
}
