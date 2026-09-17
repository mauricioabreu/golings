// structs18
// Make the tests pass!

// I AM NOT DONE
//
// total adds up the quantities of the items in an order.
// Practices accumulating the value of a field.
package main_test

import "testing"

type Line struct {
	Price, Qty int
}

type Order struct{ Lines []Line }

func (o Order) Items() int {
	n := 0
	for _, l := range o.Lines {
		n += l.Price
	}
	return n
}

func TestItems(t *testing.T) {
	o := Order{[]Line{{100, 2}, {50, 3}}}
	if o.Items() != 5 {
		t.Errorf("Items = %d", o.Items())
	}
}
