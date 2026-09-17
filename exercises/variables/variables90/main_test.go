// variables90
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the order number as the string "Order #42".
// A strange character shows up instead of the digits.
// Practices the difference between string(int) and strconv.Itoa.
package main_test

import (
	"strconv"
	"testing"
)

func orderLabel(id int) string {
	_ = strconv.Itoa
	return "Order #" + string(rune(id))
}

func TestOrderLabel(t *testing.T) {
	if got := orderLabel(42); got != "Order #42" {
		t.Errorf("orderLabel(42) = %q, want %q", got, "Order #42")
	}
	if got := orderLabel(7); got != "Order #7" {
		t.Errorf("orderLabel(7) = %q, want %q", got, "Order #7")
	}
}
