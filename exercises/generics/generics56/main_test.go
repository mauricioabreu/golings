// generics56
// Make the tests pass!

// I AM NOT DONE
//
// Filter returns a []T, and the result loses the methods of the named type Names.
// The code does not compile: a []string has no Join method.
// The S ~[]E pattern keeps the slice type.
package main_test

import (
	"strings"
	"testing"
)

type Names []string

func (n Names) Join() string { return strings.Join(n, ",") }

func Filter[E any](s []E, keep func(E) bool) []E {
	var out []E
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func TestFilterNames(t *testing.T) {
	names := Names{"ann", "", "bob"}
	got := Filter(names, func(s string) bool { return s != "" }).Join()
	if got != "ann,bob" {
		t.Errorf("got %q", got)
	}
}
