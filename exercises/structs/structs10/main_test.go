// structs10
// Make the tests pass!

// I AM NOT DONE
//
// Service embeds Counter by value, and Inc is declared on the pointer.
// A Service value does not satisfy the interface, while a pointer does.
// The *T methods of an embedded T are only promoted to *Outer.
package main_test

import "testing"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

type Service struct {
	Counter
	Name string
}

type Incer interface{ Inc() }

func bump(i Incer, times int) {
	for k := 0; k < times; k++ {
		i.Inc()
	}
}

func run() int {
	s := Service{Name: "api"}
	bump(s, 3)
	return s.n
}

func TestRun(t *testing.T) {
	if got := run(); got != 3 {
		t.Errorf("run = %d, want 3", got)
	}
}
