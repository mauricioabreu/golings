// structs46
// Make the tests pass!

// I AM NOT DONE
//
// Order.Total adds up the lines and applies a discount in percent (0-100).
// Practices methods that use nested structs.
package main_test

import "testing"

type Line struct{ Price, Qty int }

type Order struct {
	Lines    []Line
	Discount int
}

func (o Order) Total() int {
	total := 0
	for _, l := range o.Lines {
		total += l.Price
	}
	return total - o.Discount
}

func TestOrderTotal(t *testing.T) {
	o := Order{[]Line{{100, 2}, {50, 4}}, 10}
	if got := o.Total(); got != 360 {
		t.Errorf("Total = %d, want 360", got)
	}
}
