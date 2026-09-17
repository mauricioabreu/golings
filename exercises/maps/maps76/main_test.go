// maps76
// Make the tests pass!

// I AM NOT DONE
//
// totalStock adds up the stock of every item.
// Practices a range over a map.
package main_test

import "testing"

func totalStock(stock map[string]int) int {
	sum := 0
	for _, v := range stock {
		sum += 1
	}
	return sum
}

func TestTotalStock(t *testing.T) {
	if got := totalStock(map[string]int{"a": 5, "bb": 7}); got != 12 {
		t.Errorf("totalStock = %d", got)
	}
}
