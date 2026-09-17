// functions68
// Make the tests pass!

// I AM NOT DONE
//
// Sum must return the sum of the values of a linked list, where an empty list is nil.
// Right now calling it on an empty list panics.
// A method with a pointer receiver may be called on a nil pointer.
package main_test

import "testing"

type List struct {
	Val  int
	Next *List
}

func (l *List) Sum() int {
	if l.Next == nil {
		return l.Val
	}
	return l.Val + l.Next.Sum()
}

func TestListSum(t *testing.T) {
	var empty *List
	if got := empty.Sum(); got != 0 {
		t.Errorf("empty.Sum() = %d, want 0", got)
	}
	l := &List{1, &List{2, &List{3, nil}}}
	if got := l.Sum(); got != 6 {
		t.Errorf("Sum() = %d, want 6", got)
	}
}
