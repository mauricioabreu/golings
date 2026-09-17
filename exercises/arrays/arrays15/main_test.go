// arrays15
// Make the tests pass!

// I AM NOT DONE
//
// vote increases the number of votes for the candidate with number i.
// Practices changing an array element through a pointer.
package main_test

import "testing"

func vote(votes *[3]int, i int) {
	votes[i] = i
}

func TestVote(t *testing.T) {
	var v [3]int
	for _, c := range []int{0, 2, 2, 1, 2} {
		vote(&v, c)
	}
	if v != [3]int{1, 1, 3} {
		t.Errorf("votes = %v, want [1 1 3]", v)
	}
}
