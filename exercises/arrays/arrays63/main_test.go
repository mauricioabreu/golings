// arrays63
// Make the tests pass!

// I AM NOT DONE
//
// load copies a packet into a reusable [8]byte buffer and returns the length.
// When a short packet follows a long one, stray bytes are left in the buffer.
// copy does not clear the rest of the array.
package main_test

import "testing"

func load(buf *[8]byte, packet []byte) int {
	n := copy(buf[:], packet)
	return n
}

func TestLoad(t *testing.T) {
	var buf [8]byte
	load(&buf, []byte("abcdefgh"))
	n := load(&buf, []byte("xy"))
	if n != 2 || buf != [8]byte{'x', 'y'} {
		t.Errorf("buf = %q (n=%d), want \"xy\" followed by zeros", buf, n)
	}
}
