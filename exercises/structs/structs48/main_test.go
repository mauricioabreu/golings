// structs48
// Make the tests pass!

// I AM NOT DONE
//
// Counter has to satisfy the Incrementer interface. The code does not compile:
// Inc is declared on the pointer while a value is put into the interface.
// The method set of T does not hold the methods of *T.
package main_test

import "testing"

type Incrementer interface{ Inc() int }

type Counter struct{ n int }

func (c *Counter) Inc() int {
	c.n++
	return c.n
}

func newIncrementer() Incrementer {
	return Counter{}
}

func TestIncrementer(t *testing.T) {
	i := newIncrementer()
	i.Inc()
	if got := i.Inc(); got != 2 {
		t.Errorf("Inc = %d, want 2", got)
	}
}
