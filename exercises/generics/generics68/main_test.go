// generics68
// Make the tests pass!

// I AM NOT DONE
//
// Chunk splits a slice of any type into parts of size n.
// Practices generic slices of slices.
package main_test

import (
	"reflect"
	"testing"
)

func Chunk[T any](s []T, n int) [][]T {
	var out [][]T
	for i := 0; i < len(s); i += n {
		out = append(out, s[i:i+n])
	}
	return out
}

func TestChunk(t *testing.T) {
	got := Chunk([]string{"a", "b", "c"}, 2)
	if !reflect.DeepEqual(got, [][]string{{"a", "b"}, {"c"}}) {
		t.Errorf("Chunk = %v", got)
	}
}
