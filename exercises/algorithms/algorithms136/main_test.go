// algorithms136
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: BFS over an implicit graph. Find the length of the shortest
// transformation chain beginWord → endWord, where each step changes one
// letter and every intermediate word is in the dictionary. If there is no
// chain — 0. The length is counted in words, including the first and last.
// Expected asymptotics: O(N·L²) time, O(N·L) space.
package main_test

import "testing"

func ladderLength(begin, end string, words []string) int {
	return 0
}

func TestLadderLength(t *testing.T) {
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	if got := ladderLength("hit", "cog", words); got != 5 {
		t.Errorf("ladderLength = %d, want 5", got)
	}
	if got := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}); got != 0 {
		t.Errorf("missing end word should give 0, got %d", got)
	}
	if got := ladderLength("a", "c", []string{"a", "b", "c"}); got != 2 {
		t.Errorf("single letter ladder = %d, want 2", got)
	}
}
