// slices84
// Make the tests pass!

// I AM NOT DONE
//
// buffer builds an empty slice with a capacity of 16.
// Practices make([]T, len, cap).
package main_test

import "testing"

func buffer() []byte {
	return make([]byte, 16)
}

func TestBuffer(t *testing.T) {
	b := buffer()
	if len(b) != 0 || cap(b) != 16 {
		t.Errorf("len=%d cap=%d", len(b), cap(b))
	}
}
