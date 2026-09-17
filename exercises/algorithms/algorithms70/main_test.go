// algorithms70
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: backtracking with sets of occupied lines. Count in how many ways
// n queens can be placed on an n×n board so that none of them attack each
// other.
// Expected asymptotics: O(n!) time, O(n) space.
package main_test

import "testing"

func totalNQueens(n int) int {
	return 0
}

func TestTotalNQueens(t *testing.T) {
	cases := map[int]int{1: 1, 2: 0, 3: 0, 4: 2, 5: 10, 6: 4, 8: 92}
	for n, want := range cases {
		if got := totalNQueens(n); got != want {
			t.Errorf("totalNQueens(%d) = %d, want %d", n, got, want)
		}
	}
}
