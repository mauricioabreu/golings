// slices25
// Make the tests pass!

// I AM NOT DONE
//
// upperFirstField uppercases the first field of a line and returns
// a copy of it; the original line must not change.
// The results of bytes.Split refer to the original slice.
package main_test

import (
	"bytes"
	"testing"
)

func upperFirstField(line []byte) []byte {
	f := bytes.Split(line, []byte(","))[0]
	for i, c := range f {
		if c >= 'a' && c <= 'z' {
			f[i] = c - 32
		}
	}
	return f
}

func TestUpperFirstField(t *testing.T) {
	line := []byte("ann,bob")
	got := upperFirstField(line)
	if string(got) != "ANN" {
		t.Errorf("upperFirstField = %q", got)
	}
	if string(line) != "ann,bob" {
		t.Errorf("line modified: %q", line)
	}
}
