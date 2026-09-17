// variables61
// Make the tests pass!

// I AM NOT DONE
//
// reverse must reverse a string containing any Unicode characters.
// "λόγος" comes out as garbage.
// A string is UTF-8 bytes, while a character is a rune.
package main_test

import "testing"

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func TestReverse(t *testing.T) {
	cases := map[string]string{"go": "og", "λόγος": "ςογόλ", "": "", "αb": "bα"}
	for in, want := range cases {
		if got := reverse(in); got != want {
			t.Errorf("reverse(%q) = %q, want %q", in, got, want)
		}
	}
}
