// maps33
// Make the tests pass!

// I AM NOT DONE
//
// byLength groups the words by their length.
// Practices map[int][]string and appending to a value.
package main_test

import (
	"reflect"
	"testing"
)

func byLength(words []string) map[int][]string {
	g := map[int][]string{}
	for _, w := range words {
		g[len(g)] = append(g[len(g)], w)
	}
	return g
}

func TestByLength(t *testing.T) {
	g := byLength([]string{"go", "c", "js", "rust"})
	if !reflect.DeepEqual(g[2], []string{"go", "js"}) || len(g[1]) != 1 || len(g[4]) != 1 {
		t.Errorf("byLength = %v", g)
	}
}
