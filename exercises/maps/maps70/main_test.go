// maps70
// Make the tests pass!

// I AM NOT DONE
//
// sortedKeys returns the keys of a map in alphabetical order.
// Practices collecting keys and sort.Strings.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return keys
}

func TestSortedKeys(t *testing.T) {
	if got := sortedKeys(map[string]int{"b": 1, "c": 2, "a": 3}); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("sortedKeys = %v", got)
	}
}
