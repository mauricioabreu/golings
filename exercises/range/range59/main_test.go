// range59
// Make the tests pass!

// I AM NOT DONE
//
// wordFreq counts words ignoring case, splitting the text on any non-letter character.
// Practices a range over runes while building up the current word.
package main_test

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func wordFreq(text string) map[string]int {
	m := map[string]int{}
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 0 {
			m[cur.String()]++
			cur.Reset()
		}
	}
	for _, r := range text {
		if r == ' ' {
			flush()
			continue
		}
		cur.WriteRune(r)
	}
	return m
}

func TestWordFreq(t *testing.T) {
	_ = unicode.IsLetter
	got := wordFreq("Go, go! Καλα και ΚΑΛΑ.")
	want := map[string]int{"go": 2, "καλα": 2, "και": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wordFreq = %v", got)
	}
}
