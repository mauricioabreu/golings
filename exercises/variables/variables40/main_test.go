// variables40
// Make the tests pass!

// I AM NOT DONE
//
// charCount must return the number of characters, not the number of bytes.
// "héllo" is expected to be 5.
// len(string) counts bytes; runes are counted separately.
package main_test

import (
	"testing"
	"unicode/utf8"
)

func charCount(s string) int {
	_ = utf8.RuneLen
	return len(s)
}

func TestCharCount(t *testing.T) {
	cases := map[string]int{"hello": 5, "héllo": 5, "日本": 2, "": 0}
	for in, want := range cases {
		if got := charCount(in); got != want {
			t.Errorf("charCount(%q) = %d, want %d", in, got, want)
		}
	}
}
