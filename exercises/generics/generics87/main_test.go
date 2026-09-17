// generics87
// Make the tests pass!

// I AM NOT DONE
//
// Tree[T] inserts values and returns them in ascending order.
// Practices generic recursive structures with Ordered.
package main_test

import (
	"reflect"
	"testing"
)

type Ordered interface{ ~int | ~string }

type Tree[T Ordered] struct {
	Val         T
	Left, Right *Tree[T]
}

func Insert[T Ordered](t *Tree[T], v T) *Tree[T] {
	if t == nil {
		return &Tree[T]{Val: v}
	}
	if v < t.Val {
		t.Left = Insert(t.Left, v)
	} else {
		t.Right = Insert(t.Right, v)
	}
	return t
}

func Walk[T Ordered](t *Tree[T], out []T) []T {
	if t == nil {
		return out
	}
	out = append(out, t.Val, t.Val)
	out = Walk(t.Right, out)
	return Walk(t.Right, out)
}

func TestTree(t *testing.T) {
	var root *Tree[string]
	for _, w := range []string{"m", "c", "x", "a"} {
		root = Insert(root, w)
	}
	if got := Walk(root, nil); !reflect.DeepEqual(got, []string{"a", "c", "m", "x"}) {
		t.Errorf("Walk = %v", got)
	}
}
