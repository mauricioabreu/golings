// switch49
// Make the tests pass!

// I AM NOT DONE
//
// The order status is declared with iota. A new order with no status set
// counts as "paid", although the status is simply unset.
// The zero value of an enum should mean "unset".
package main_test

import "testing"

type Status int

const (
	Paid Status = iota
	Shipped
)

type Order struct {
	ID     int
	Status Status
}

func label(o Order) string {
	switch o.Status {
	case Paid:
		return "paid"
	case Shipped:
		return "shipped"
	default:
		return "unset"
	}
}

func TestLabel(t *testing.T) {
	if got := label(Order{ID: 1}); got != "unset" {
		t.Errorf("label(new order) = %s, want unset", got)
	}
	if got := label(Order{Status: Paid}); got != "paid" {
		t.Errorf("label(paid) = %s", got)
	}
	if got := label(Order{Status: Shipped}); got != "shipped" {
		t.Errorf("label(shipped) = %s", got)
	}
}
