// anonymous_functions68
// Make the tests pass!

// I AM NOT DONE
//
// The literal has to add a bonus to the total score, and its parameter carries
// the same name as the outer variable, so the changes are lost.
// A parameter of a literal shadows a captured variable.
package main_test

import "testing"

func play(bonuses []int) int {
	score := 0
	addBonus := func(score int) {
		score += score
	}
	for _, b := range bonuses {
		addBonus(b)
	}
	return score
}

func TestPlay(t *testing.T) {
	if got := play([]int{5, 10}); got != 15 {
		t.Errorf("play = %d, want 15", got)
	}
}
