// generics8
// Make the tests pass!

// I AM NOT DONE
//
// Ring[T] keeps the last n values; Items returns them from the oldest to the newest.
// Practices a generic type with a fixed capacity.
package main_test

import (
	"reflect"
	"testing"
)

type Ring[T any] struct {
	buf        []T
	pos, count int
}

func NewRing[T any](n int) *Ring[T] { return &Ring[T]{buf: make([]T, n)} }

func (r *Ring[T]) Add(v T) {
	r.buf[r.pos] = v
	r.pos++
	r.count++
}

func (r *Ring[T]) Items() []T {
	n := len(r.buf)
	out := make([]T, 0, r.count)
	start := (r.pos - r.count + n) % n
	for i := 0; i < r.count; i++ {
		out = append(out, r.buf[(start+i)%n])
	}
	return out
}

func TestRing(t *testing.T) {
	r := NewRing[string](2)
	for _, s := range []string{"a", "b", "c"} {
		r.Add(s)
	}
	if !reflect.DeepEqual(r.Items(), []string{"b", "c"}) {
		t.Errorf("Items = %v", r.Items())
	}
}
