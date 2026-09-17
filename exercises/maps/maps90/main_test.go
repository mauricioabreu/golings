// maps90
// Make the tests pass!

// I AM NOT DONE
//
// byValue groups the keys by value, the values possibly repeating,
// with the key lists sorted.
// Practices map[V][]K and sorting the values of a map.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func byValue(m map[string]int) map[int][]string {
	out := map[int][]string{}
	for k, v := range m {
		out[v] = []string{k}
	}
	return out
}

func TestByValue(t *testing.T) {
	got := byValue(map[string]int{"a": 1, "b": 2, "c": 1, "d": 1})
	want := map[int][]string{1: {"a", "c", "d"}, 2: {"b"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("byValue = %v", got)
	}
}
