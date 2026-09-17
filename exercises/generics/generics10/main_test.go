// generics10
// Make the tests pass!

// I AM NOT DONE
//
// Filter keeps the elements satisfying a predicate.
// Practices generic higher order functions.
package main_test

import (
	"reflect"
	"testing"
)

func Filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if !keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func TestFilter(t *testing.T) {
	got := Filter([]string{"go", "", "c"}, func(s string) bool { return s != "" })
	if !reflect.DeepEqual(got, []string{"go", "c"}) {
		t.Errorf("Filter = %v", got)
	}
}
