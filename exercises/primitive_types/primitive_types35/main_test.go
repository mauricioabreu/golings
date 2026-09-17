// primitive_types35
// Make the tests pass!

// I AM NOT DONE
//
// cString turns a fixed size buffer holding a C string, terminated by a zero
// byte, into a Go string. Right now the zero bytes end up in the string.
// string(buf) copies every byte, \x00 included.
package main_test

import (
	"bytes"
	"testing"
)

func cString(buf []byte) string {
	return string(buf)
}

func TestCString(t *testing.T) {
	_ = bytes.IndexByte
	buf := make([]byte, 16)
	copy(buf, "gopher")
	if got := cString(buf); got != "gopher" {
		t.Errorf("cString = %q, want gopher", got)
	}
	if got := cString([]byte("full")); got != "full" {
		t.Errorf("cString(no terminator) = %q", got)
	}
}
