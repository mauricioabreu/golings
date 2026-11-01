// methods3
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c Counter) Value() int {
	return c.count
}

// incrementAll adds one to every counter in the slice.
func incrementAll(counters []Counter) {
	// TODO: Increment the counters in the slice, not copies of them.
	for _, c := range counters {
		c.Increment()
	}
}

func TestIncrementAll(t *testing.T) {
	counters := make([]Counter, 3)

	incrementAll(counters)

	for i, c := range counters {
		if c.Value() != 1 {
			t.Errorf("counter %d should be 1, but is %d", i, c.Value())
		}
	}
}
