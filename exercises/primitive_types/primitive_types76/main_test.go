// primitive_types76
// Make the tests pass!

// I AM NOT DONE
//
// hexID must return a number as lowercase hexadecimal digits.
// Practices the %x verb.
package main_test

import (
	"fmt"
	"testing"
)

func hexID(n int) string {
	return fmt.Sprintf("%X", n)
}

func TestHexID(t *testing.T) {
	cases := map[int]string{255: "ff", 3054: "bee", 16: "10"}
	for in, want := range cases {
		if got := hexID(in); got != want {
			t.Errorf("hexID(%d) = %s, want %s", in, got, want)
		}
	}
}
