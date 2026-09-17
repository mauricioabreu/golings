// functions22
// Make the tests pass!

// I AM NOT DONE
//
// snapshotter returns a method value that must report the CURRENT
// value of the counter.
// A method value with a value receiver copies the receiver when it is evaluated.
package main_test

import "testing"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

func (c Counter) Value() int { return c.n }

func TestMethodValue(t *testing.T) {
	c := &Counter{}
	value := c.Value
	c.Inc()
	c.Inc()
	if got := value(); got != 2 {
		t.Errorf("value() = %d, want 2", got)
	}
}
