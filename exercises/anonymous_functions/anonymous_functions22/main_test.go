// anonymous_functions22
// Make the tests pass!

// I AM NOT DONE
//
// bytes.IndexFunc expects a func(rune) bool. The literal is declared with a byte parameter
// and the code does not compile.
// The parameter types of a literal have to match exactly.
package main_test

import (
	"bytes"
	"testing"
)

func firstSpace(b []byte) int {
	return bytes.IndexFunc(b, func(c byte) bool { return c == ' ' })
}

func TestFirstSpace(t *testing.T) {
	if firstSpace([]byte("go lang")) != 2 {
		t.Errorf("firstSpace = %d", firstSpace([]byte("go lang")))
	}
}
