// slices26
// Make the tests pass!

// I AM NOT DONE
//
// appendText appends a string to a byte buffer.
// Practices append([]byte, string...).
package main_test

import "testing"

func appendText(buf []byte, s string) []byte {
	return append(buf, s[0])
}

func TestAppendText(t *testing.T) {
	if got := appendText([]byte("go"), "pher"); string(got) != "gopher" {
		t.Errorf("appendText = %q", got)
	}
}
