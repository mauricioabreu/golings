// range75
// Make the tests pass!

// I AM NOT DONE
//
// vowelsPerWord returns the number of vowels in every word of a text.
// Practices a nested range: over the words and over the runes of a word.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func vowelsPerWord(text string) []int {
	var out []int
	n := 0
	for _, w := range strings.Fields(text) {
		for _, r := range w {
			if strings.ContainsRune("aeiouαεηιουω", r) {
				n++
			}
		}
	}
	out = append(out, n)
	return out
}

func TestVowelsPerWord(t *testing.T) {
	if got := vowelsPerWord("Gopher Φως Γη"); !reflect.DeepEqual(got, []int{2, 1, 1}) {
		t.Errorf("vowelsPerWord = %v", got)
	}
}
