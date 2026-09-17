// algorithms37
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: trie and backtracking. Find all words from the dictionary that
// can be built on the board by moving to neighboring cells (without
// reusing a cell). The answer is sorted alphabetically.
// Expected asymptotics: O(r·c·4^L) time, O(total word length) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func findWords(board [][]byte, words []string) []string {
	return nil
}

func TestFindWords(t *testing.T) {
	_ = sort.Strings
	board := [][]byte{
		[]byte("oaan"),
		[]byte("etae"),
		[]byte("ihkr"),
		[]byte("iflv"),
	}
	got := findWords(board, []string{"oath", "pea", "eat", "rain"})
	if !reflect.DeepEqual(got, []string{"eat", "oath"}) {
		t.Errorf("findWords = %v", got)
	}
	if got := findWords([][]byte{[]byte("ab")}, []string{"ba", "ab", "abc"}); !reflect.DeepEqual(got, []string{"ab", "ba"}) {
		t.Errorf("small board = %v", got)
	}
	if got := findWords(nil, []string{"a"}); len(got) != 0 {
		t.Errorf("empty board = %v", got)
	}
}
