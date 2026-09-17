// anonymous_functions23
// Make the tests pass!

// I AM NOT DONE
//
// groupBy groups strings by a key computed by a literal.
// byFirstLetter groups by the first letter in lower case.
// Practices a key function as a parameter.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func groupBy(items []string, key func(string) string) map[string][]string {
	out := map[string][]string{}
	for _, it := range items {
		out[it] = append(out[it], key(it))
	}
	return out
}

func byFirstLetter(words []string) map[string][]string {
	return groupBy(words, func(w string) string { return strings.ToLower(w[:1]) })
}

func TestByFirstLetter(t *testing.T) {
	got := byFirstLetter([]string{"Apple", "avocado", "banana"})
	if !reflect.DeepEqual(got, map[string][]string{"a": {"Apple", "avocado"}, "b": {"banana"}}) {
		t.Errorf("byFirstLetter = %v", got)
	}
}
