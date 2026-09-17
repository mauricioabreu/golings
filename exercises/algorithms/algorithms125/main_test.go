// algorithms125
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a dictionary and dynamic programming over a string. Find the
// words that are entirely composed of two or more other words of the
// dictionary. The answer is sorted alphabetically.
// Expected asymptotics: O(n·L²) time, O(n·L) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func concatenatedWords(words []string) []string {
	return nil
}

func TestConcatenatedWords(t *testing.T) {
	_ = sort.Strings
	got := concatenatedWords([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"})
	want := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("concatenatedWords = %v, want %v", got, want)
	}
	if got := concatenatedWords([]string{"cat", "dog"}); len(got) != 0 {
		t.Errorf("no concatenated words expected, got %v", got)
	}
	if got := concatenatedWords([]string{"a", "aa", "aaa"}); !reflect.DeepEqual(got, []string{"aa", "aaa"}) {
		t.Errorf("repeated letters = %v", got)
	}
}
