// range79
// Make the tests pass!

// I AM NOT DONE
//
// validRunes counts the characters of a string along with the invalid bytes.
// A range replaces invalid bytes with U+FFFD, and they have to be told apart.
// On a decoding error a range yields utf8.RuneError.
package main_test

import (
	"testing"
	"unicode/utf8"
)

func validRunes(s string) (ok, bad int) {
	for range s {
		ok++
	}
	return
}

func TestValidRunes(t *testing.T) {
	_ = utf8.RuneError
	ok, bad := validRunes("a\xffλ�\xfe")
	if ok != 3 || bad != 2 {
		t.Errorf("validRunes = %d, %d; want 3, 2", ok, bad)
	}
}
