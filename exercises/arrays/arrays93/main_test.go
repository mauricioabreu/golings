// arrays93
// Make the tests pass!

// I AM NOT DONE
//
// visits counts visits to cells using a [2]int coordinate as the key.
// Practices arrays as map keys.
package main_test

import "testing"

func visits(path [][2]int) map[[2]int]int {
	m := map[[2]int]int{}
	for _, p := range path {
		m[p] = 1
	}
	return m
}

func TestVisits(t *testing.T) {
	m := visits([][2]int{{0, 0}, {1, 0}, {0, 0}})
	if m[[2]int{0, 0}] != 2 || m[[2]int{1, 0}] != 1 {
		t.Errorf("visits = %v", m)
	}
}
