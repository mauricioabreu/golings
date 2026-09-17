// range69
// Make the tests pass!

// I AM NOT DONE
//
// totalScore must add up the scores of the players from a map of player number to score.
// The sum comes out as the sum of the numbers.
// A range over a map with a single variable yields the keys.
package main_test

import "testing"

func totalScore(scores map[int]int) int {
	sum := 0
	for v := range scores {
		sum += v
	}
	return sum
}

func TestTotalScore(t *testing.T) {
	if got := totalScore(map[int]int{1: 100, 2: 250, 3: 50}); got != 400 {
		t.Errorf("totalScore = %d, want 400", got)
	}
}
