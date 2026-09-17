// arrays62
// Make the tests pass!

// I AM NOT DONE
//
// prefix builds an array of prefix sums: p[i] = a[0] + ... + a[i-1], with p[0] = 0.
// rangeSum uses it for the sum over the range [l, r).
// Practices an array of length N+1 and working with indexes.
package main_test

import "testing"

func prefix(a [5]int) [6]int {
	var p [6]int
	for i, v := range a {
		p[i] = p[i] + v
	}
	return p
}

func rangeSum(p [6]int, l, r int) int {
	return p[r] - p[l-1]
}

func TestPrefix(t *testing.T) {
	p := prefix([5]int{3, 1, 4, 1, 5})
	if p != [6]int{0, 3, 4, 8, 9, 14} {
		t.Errorf("prefix = %v", p)
	}
	if got := rangeSum(p, 1, 4); got != 6 {
		t.Errorf("rangeSum(1, 4) = %d, want 6", got)
	}
	if got := rangeSum(p, 0, 5); got != 14 {
		t.Errorf("rangeSum(0, 5) = %d, want 14", got)
	}
}
