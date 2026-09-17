// maps36
// Make the tests pass!

// I AM NOT DONE
//
// groupAnagrams groups the words that are anagrams of each other.
// Practices a computed map key.
package main_test

import (
	"sort"
	"testing"
)

func sortedLetters(w string) string {
	b := []byte(w)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func groupAnagrams(words []string) map[string][]string {
	groups := map[string][]string{}
	for _, w := range words {
		groups[w] = append(groups[w], w)
	}
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	g := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if len(g) != 3 || len(g["aet"]) != 3 || len(g["ant"]) != 2 || len(g["abt"]) != 1 {
		t.Errorf("groupAnagrams = %v", g)
	}
}
