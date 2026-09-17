// arrays43
// Make the tests pass!

// I AM NOT DONE
//
// emptyScores must return an array of three zeros.
// Practices the zero value of an array.
package main_test

import "testing"

func emptyScores() [3]int {
	var scores [3]int
	scores[1] = 1
	return scores
}

func TestEmptyScores(t *testing.T) {
	if got := emptyScores(); got != [3]int{} {
		t.Errorf("emptyScores = %v, want [0 0 0]", got)
	}
}
