// maps74
// Make the tests pass!

// I AM NOT DONE
//
// adjust changes the stock of an item by delta. The stock must not go negative,
// which is an error, and an item with no stock left is removed from the map.
// Practices checks before a change, and delete.
package main_test

import (
	"errors"
	"testing"
)

func adjust(stock map[string]int, item string, delta int) error {
	n := stock[item] + delta
	stock[item] = n
	return nil
}

func TestAdjust(t *testing.T) {
	_ = errors.New
	s := map[string]int{"pen": 2}
	if err := adjust(s, "pen", -2); err != nil || len(s) != 0 {
		t.Errorf("adjust to zero: %v, %v", err, s)
	}
	if err := adjust(s, "cup", -1); err == nil || len(s) != 0 {
		t.Errorf("adjust below zero: %v, %v", err, s)
	}
	if err := adjust(s, "cup", 3); err != nil || s["cup"] != 3 {
		t.Errorf("adjust up: %v, %v", err, s)
	}
}
