// generics22
// Make the tests pass!

// I AM NOT DONE
//
// Remove drops every element equal to x from a slice. The code does not compile.
// any does not guarantee that == works.
package main_test

import (
	"reflect"
	"testing"
)

func Remove[T any](s []T, x T) []T {
	var out []T
	for _, v := range s {
		if v != x {
			out = append(out, v)
		}
	}
	return out
}

func TestRemove(t *testing.T) {
	if got := Remove([]string{"a", "b", "a"}, "a"); !reflect.DeepEqual(got, []string{"b"}) {
		t.Errorf("Remove = %v", got)
	}
}
