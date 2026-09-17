// functions63
// Make the tests pass!

// I AM NOT DONE
//
// bumpAll increments every counter through the Incrementer interface.
// The methods are called, yet the values stay put.
// A method with a value receiver works on a copy.
package main_test

import "testing"

type Incrementer interface {
	Inc()
	Value() int
}

type Counter struct{ n int }

func (c Counter) Inc()       { c.n++ }
func (c Counter) Value() int { return c.n }

func bumpAll(items []Incrementer, times int) {
	for _, it := range items {
		for i := 0; i < times; i++ {
			it.Inc()
		}
	}
}

func TestBumpAll(t *testing.T) {
	a, b := &Counter{}, &Counter{n: 10}
	bumpAll([]Incrementer{a, b}, 3)
	if a.Value() != 3 || b.Value() != 13 {
		t.Errorf("values = %d, %d; want 3, 13", a.Value(), b.Value())
	}
}
