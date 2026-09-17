// anonymous_functions80
// Make the tests pass!

// I AM NOT DONE
//
// byLength sorts the words by length with a literal in sort.Slice.
// Practices a function literal as an argument.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func byLength(words []string) {
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
}

func TestByLength(t *testing.T) {
	w := []string{"banana", "fig", "apple"}
	byLength(w)
	if !reflect.DeepEqual(w, []string{"fig", "apple", "banana"}) {
		t.Errorf("byLength = %v", w)
	}
}
