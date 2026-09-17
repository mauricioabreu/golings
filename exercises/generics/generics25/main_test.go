// generics25
// Make the tests pass!

// I AM NOT DONE
//
// NewAll builds n resettable objects of type T. For T = *Buffer
// the zero value is nil, and calling Reset panics.
// The zero value of a pointer type is not a ready object.
package main_test

import "testing"

type Buffer struct{ data []byte }

func (b *Buffer) Reset() { b.data = b.data[:0] }

type Resetter interface{ Reset() }

func NewAll[T Resetter](n int) []T {
	out := make([]T, n)
	for i := range out {
		out[i].Reset()
	}
	return out
}

func buffers() []*Buffer { return NewAll[*Buffer](2) }

func TestBuffers(t *testing.T) {
	bs := buffers()
	if len(bs) != 2 || bs[0] == nil || bs[1] == nil || bs[0] == bs[1] {
		t.Errorf("buffers = %v", bs)
	}
}
