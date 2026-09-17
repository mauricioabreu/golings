// generics76
// Make the tests pass!

// I AM NOT DONE
//
// Node[T] refers to itself, and without a type argument. The code does not compile.
// Inside a generic type a reference to it needs the [T].
package main_test

import "testing"

type Node[T any] struct {
	Val  T
	Next *Node
}

func Len[T any](n *Node[T]) int {
	c := 0
	for ; n != nil; n = n.Next {
		c++
	}
	return c
}

func TestLen(t *testing.T) {
	list := &Node[string]{"a", &Node[string]{"b", nil}}
	if Len(list) != 2 {
		t.Errorf("Len = %d", Len(list))
	}
}
