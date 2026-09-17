// slices10
// Make the tests pass!

// I AM NOT DONE
//
// firstLine returns the first line of a read buffer. The buffer is reused later,
// and the string that was returned goes bad.
// A subslice refers to the same array as the buffer.
package main_test

import (
	"bytes"
	"testing"
)

func firstLine(buf []byte) []byte {
	i := bytes.IndexByte(buf, '\n')
	if i < 0 {
		i = len(buf)
	}
	return buf[:i]
}

func TestFirstLine(t *testing.T) {
	buf := []byte("hello\nworld")
	line := firstLine(buf)
	copy(buf, "XXXXX")
	if string(line) != "hello" {
		t.Errorf("line = %q, want hello", line)
	}
}
