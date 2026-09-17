// range16
// Make the tests pass!

// I AM NOT DONE
//
// runeWidths returns the width in bytes of every character, decoding the string by hand.
// The loop moves one byte at a time and sees characters that are not there.
// utf8.DecodeRuneInString returns the width to move forward by.
package main_test

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func runeWidths(s string) []int {
	var out []int
	for i := 0; i < len(s); i++ {
		_, w := utf8.DecodeRuneInString(s[i:])
		out = append(out, w)
	}
	return out
}

func TestRuneWidths(t *testing.T) {
	if got := runeWidths("aλ€"); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("runeWidths = %v, want [1 2 3]", got)
	}
}
