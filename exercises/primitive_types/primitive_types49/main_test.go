// primitive_types49
// Make the tests pass!

// I AM NOT DONE
//
// countRunes must return the number of characters, or an error when the string holds
// invalid UTF-8 bytes. Right now the junk bytes are counted as characters.
// A range over an invalid string yields utf8.RuneError.
package main_test

import (
	"errors"
	"testing"
	"unicode/utf8"
)

func countRunes(s string) (int, error) {
	return utf8.RuneCountInString(s), nil
}

func TestCountRunes(t *testing.T) {
	_ = errors.New
	if n, err := countRunes("héllo"); err != nil || n != 5 {
		t.Errorf("countRunes(héllo) = %d, %v", n, err)
	}
	if _, err := countRunes("bad\xff\xfe"); err == nil {
		t.Errorf("countRunes with invalid bytes should fail")
	}
}
