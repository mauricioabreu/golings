// primitive_types99
// Make the tests pass!

// I AM NOT DONE
//
// rawString must build a string out of raw bytes rather than text.
// For bytes above 127 the length of the result does not match the number of bytes.
// string(byte(200)) encodes the rune U+00C8 in UTF-8, which takes two bytes.
package main_test

import "testing"

func rawString(data []byte) string {
	s := ""
	for _, b := range data {
		s += string(rune(b))
	}
	return s
}

func TestRawString(t *testing.T) {
	data := []byte{0x41, 0xC8, 0xFF}
	got := rawString(data)
	if len(got) != 3 || got[1] != 0xC8 || got[2] != 0xFF {
		t.Errorf("rawString = % x, want 41 c8 ff", got)
	}
}
