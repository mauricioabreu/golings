// functions54
// Make the tests pass!

// I AM NOT DONE
//
// splitWords must split a string on any non-letter character.
// Practices strings.FieldsFunc and predicate functions.
package main_test

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func splitWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsLetter(r)
	})
}

func TestSplitWords(t *testing.T) {
	got := splitWords("go,is;fun!  yes")
	if want := []string{"go", "is", "fun", "yes"}; !reflect.DeepEqual(got, want) {
		t.Errorf("splitWords = %v, want %v", got, want)
	}
}
