// switch32
// Make the tests pass!

// I AM NOT DONE
//
// ratingStars turns a score of 0-100 into stars: 90+ gives 3, 70+ gives 2, 50+ gives 1, else 0.
// Practices the order of the branches in a tagless switch.
package main_test

import "testing"

func ratingStars(score int) int {
	switch {
	case score >= 90:
		return 3
	case score > 70:
		return 2
	case score >= 50:
		return 1
	}
	return 0
}

func TestRatingStars(t *testing.T) {
	cases := map[int]int{95: 3, 90: 3, 70: 2, 69: 1, 50: 1, 10: 0}
	for in, want := range cases {
		if got := ratingStars(in); got != want {
			t.Errorf("ratingStars(%d) = %d, want %d", in, got, want)
		}
	}
}
