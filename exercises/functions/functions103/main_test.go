// functions103
// Make the tests pass!

// I AM NOT DONE
//
// addItem must add an item to the caller's cart.
// The length of the caller's cart is unchanged after the call.
// A slice is passed by value: the header is copied.
package main_test

import (
	"reflect"
	"testing"
)

func addItem(cart []string, item string) {
	cart = append(cart, item)
}

func fillCart() []string {
	cart := make([]string, 0, 10)
	addItem(cart, "apple")
	addItem(cart, "milk")
	return cart
}

func TestFillCart(t *testing.T) {
	if got := fillCart(); !reflect.DeepEqual(got, []string{"apple", "milk"}) {
		t.Errorf("fillCart() = %v, want [apple milk]", got)
	}
}
