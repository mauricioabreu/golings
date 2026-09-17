// algorithms12
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: hash tables with a computed key. Group anagram words together.
// Groups are ordered by first appearance; words within a group — as in
// the input.
// Expected asymptotics: O(n·k·log k) time, O(n·k) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func groupAnagrams(words []string) [][]string {
	return nil
}

func TestGroupAnagrams(t *testing.T) {
	_ = sort.Ints
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("groupAnagrams = %v, want %v", got, want)
	}
	if got := groupAnagrams([]string{""}); !reflect.DeepEqual(got, [][]string{{""}}) {
		t.Errorf("groupAnagrams([\"\"]) = %v", got)
	}
	if got := groupAnagrams(nil); len(got) != 0 {
		t.Errorf("groupAnagrams(nil) = %v", got)
	}
}
