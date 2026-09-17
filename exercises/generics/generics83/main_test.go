// generics83
// Make the tests pass!

// I AM NOT DONE
//
// GroupBy groups the elements by a key computed by a function.
// Practices two type parameters, one of them comparable.
package main_test

import (
	"reflect"
	"testing"
)

func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	out := map[K][]T{}
	for _, v := range s {
		out[key(v)] = []T{v}
	}
	return out
}

func TestGroupBy(t *testing.T) {
	got := GroupBy([]int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 })
	if !reflect.DeepEqual(got, map[bool][]int{true: {2, 4}, false: {1, 3, 5}}) {
		t.Errorf("GroupBy = %v", got)
	}
}
