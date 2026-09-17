// generics59
// Make the tests pass!

// I AM NOT DONE
//
// Pair[K, V] holds a key and a value; Swap exchanges them.
// Practices a generic struct with two parameters.
package main_test

import "testing"

type Pair[K, V any] struct {
	Key K
	Val V
}

func Swap[K, V any](p Pair[K, V]) Pair[V, K] {
	return Pair[V, K]{Key: p.Val}
}

func TestSwap(t *testing.T) {
	s := Swap(Pair[string, int]{"age", 30})
	if s.Key != 30 || s.Val != "age" {
		t.Errorf("Swap = %+v", s)
	}
}
