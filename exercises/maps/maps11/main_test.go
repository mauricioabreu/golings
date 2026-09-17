// maps11
// Make the tests pass!

// I AM NOT DONE
//
// totals adds up the expenses per category and returns the grand total as well.
// Practices accumulating values in a map alongside a counter.
package main_test

import (
	"reflect"
	"testing"
)

type Expense struct {
	Category string
	Amount   int
}

func totals(es []Expense) (map[string]int, int) {
	m := map[string]int{}
	sum := 0
	for _, e := range es {
		m[e.Category] = e.Amount
	}
	return m, sum
}

func TestTotals(t *testing.T) {
	m, sum := totals([]Expense{{"food", 10}, {"taxi", 5}, {"food", 7}})
	if !reflect.DeepEqual(m, map[string]int{"food": 17, "taxi": 5}) || sum != 22 {
		t.Errorf("totals = %v, %d", m, sum)
	}
}
