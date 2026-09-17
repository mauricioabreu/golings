// range58
// Make the tests pass!

// I AM NOT DONE
//
// chars returns the characters of a string as a slice of strings.
// For non-ASCII text it produces byte fragments.
// s[i:i+1] takes a single byte, not a character.
package main_test

import (
	"reflect"
	"testing"
)

func chars(s string) []string {
	var out []string
	for i, r := range s {
		out = append(out, s[i:i+1])
		_ = r
	}
	return out
}

func TestChars(t *testing.T) {
	if got := chars("δα!"); !reflect.DeepEqual(got, []string{"δ", "α", "!"}) {
		t.Errorf("chars = %q", got)
	}
}
