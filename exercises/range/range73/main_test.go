// range73
// Make the tests pass!

// I AM NOT DONE
//
// longestRun returns the character with the longest run and the length of that run.
// Practices a range over a string with a current and a best run.
package main_test

import "testing"

func longestRun(s string) (rune, int) {
	var best, prev rune
	bestN, cur := 0, 0
	for _, r := range s {
		if r == prev {
			cur++
		}
		if cur > bestN {
			best, bestN = r, cur
		}
	}
	return best, bestN
}

func TestLongestRun(t *testing.T) {
	if r, n := longestRun("aabbbbccλ"); r != 'b' || n != 4 {
		t.Errorf("longestRun = %c, %d", r, n)
	}
	if r, n := longestRun("λλλ"); r != 'λ' || n != 3 {
		t.Errorf("longestRun(λλλ) = %c, %d", r, n)
	}
}
