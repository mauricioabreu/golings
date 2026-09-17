// maps29
// Make the tests pass!

// I AM NOT DONE
//
// positions returns the list of positions of every word in a text.
// Practices map[string][]int.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func positions(text string) map[string][]int {
	m := map[string][]int{}
	for _, w := range strings.Fields(text) {
		m[w] = append(m[w], len(m[w]))
	}
	return m
}

func TestPositions(t *testing.T) {
	got := positions("a b a c a")
	if !reflect.DeepEqual(got, map[string][]int{"a": {0, 2, 4}, "b": {1}, "c": {3}}) {
		t.Errorf("positions = %v", got)
	}
}
