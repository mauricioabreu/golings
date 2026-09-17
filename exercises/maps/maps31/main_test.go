// maps31
// Make the tests pass!

// I AM NOT DONE
//
// bestSeller returns the item with the highest sales, the values being distinct.
// Practices finding a maximum over a map.
package main_test

import "testing"

func bestSeller(sales map[string]int) string {
	best, max := "", -1
	for item, n := range sales {
		if n > max {
			max = n
		}
	}
	return best
}

func TestBestSeller(t *testing.T) {
	if got := bestSeller(map[string]int{"tea": 5, "coffee": 9, "juice": 2}); got != "coffee" {
		t.Errorf("bestSeller = %q", got)
	}
}
