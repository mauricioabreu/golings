// algorithms128
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: disjoint set union (union-find). Count the number of connected
// components in an undirected graph with n vertices.
// Expected asymptotics: O(E·α(V)) time, O(V) space.
package main_test

import "testing"

func countComponents(n int, edges [][2]int) int {
	return 0
}

func TestCountComponents(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  int
	}{
		{5, [][2]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
		{4, nil, 4},
		{1, nil, 1},
		{3, [][2]int{{0, 1}, {0, 1}}, 2},
	}
	for _, c := range cases {
		if got := countComponents(c.n, c.edges); got != c.want {
			t.Errorf("countComponents(%d, %v) = %d, want %d", c.n, c.edges, got, c.want)
		}
	}
}
