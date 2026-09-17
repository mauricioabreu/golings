// range14
// Make the tests pass!

// I AM NOT DONE
//
// buildIndex builds a map from a word to the numbers of the documents, with no repeats.
// Practices a nested range over a slice of strings and their words.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func buildIndex(docs []string) map[string][]int {
	idx := map[string][]int{}
	for i, d := range docs {
		for _, w := range strings.Fields(d) {
			ids := idx[w]
			idx[w] = append(ids, len(ids))
		}
	}
	return idx
}

func TestBuildIndex(t *testing.T) {
	idx := buildIndex([]string{"go is go", "rust is", "go"})
	want := map[string][]int{"go": {0, 2}, "is": {0, 1}, "rust": {1}}
	if !reflect.DeepEqual(idx, want) {
		t.Errorf("buildIndex = %v", idx)
	}
}
