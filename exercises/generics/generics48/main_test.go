// generics48
// Make the tests pass!

// I AM NOT DONE
//
// Union and Intersect work on generic sets.
// Practices generic functions over a Set[T].
package main_test

import "testing"

type Set[T comparable] map[T]struct{}

func Union[T comparable](a, b Set[T]) Set[T] {
	out := Set[T]{}
	for k := range a {
		out[k] = struct{}{}
	}
	return out
}

func Intersect[T comparable](a, b Set[T]) Set[T] {
	out := Set[T]{}
	for k := range a {
		out[k] = struct{}{}
	}
	return out
}

func TestSets(t *testing.T) {
	a := Set[int]{1: {}, 2: {}}
	b := Set[int]{2: {}, 3: {}}
	if len(Union(a, b)) != 3 {
		t.Errorf("Union = %v", Union(a, b))
	}
	i := Intersect(a, b)
	if _, ok := i[2]; !ok || len(i) != 1 {
		t.Errorf("Intersect = %v", i)
	}
}
