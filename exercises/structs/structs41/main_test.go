// structs41
// Make the tests pass!

// I AM NOT DONE
//
// Clock holds the time as minutes since midnight; Add adds minutes with wrapping,
// negative ones included, and String prints "hh:mm".
// Practices normalizing state inside methods.
package main_test

import (
	"fmt"
	"testing"
)

type Clock struct{ min int }

func (c Clock) Add(m int) Clock {
	return Clock{c.min + m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%d:%d", c.min/60, c.min%60)
}

func TestClock(t *testing.T) {
	c := Clock{23*60 + 30}
	if got := c.Add(45).String(); got != "00:15" {
		t.Errorf("23:30 + 45 = %s", got)
	}
	if got := (Clock{10}).Add(-20).String(); got != "23:50" {
		t.Errorf("00:10 - 20 = %s", got)
	}
}
