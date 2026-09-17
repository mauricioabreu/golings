// range55
// Make the tests pass!

// I AM NOT DONE
//
// lengthHistogram returns a slice where element i is the number of words of length i, in characters.
// Practices a range over words while growing the result as needed.
package main_test

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func lengthHistogram(text string) []int {
	var h []int
	for _, w := range strings.Fields(text) {
		n := len(w)
		h = append(h, n)
	}
	return h
}

func TestLengthHistogram(t *testing.T) {
	_ = utf8.RuneCountInString
	if got := lengthHistogram("α β δε ζη go"); !reflect.DeepEqual(got, []int{0, 2, 3}) {
		t.Errorf("lengthHistogram = %v", got)
	}
}
