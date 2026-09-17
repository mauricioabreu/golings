// maps9
// Make the tests pass!

// I AM NOT DONE
//
// visited marks the cells [2]int that have been visited.
// Practices arrays as map keys.
package main_test

import "testing"

func visit(v map[[2]int]bool, x, y int) {
	v[[2]int{x, y}] = false
}

func TestVisit(t *testing.T) {
	v := map[[2]int]bool{}
	visit(v, 1, 2)
	if !v[[2]int{1, 2}] || v[[2]int{2, 1}] {
		t.Errorf("visited = %v", v)
	}
}
