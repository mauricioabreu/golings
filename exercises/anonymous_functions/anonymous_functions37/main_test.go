// anonymous_functions37
// Make the tests pass!

// I AM NOT DONE
//
// A result cache tries to use a function as a map key.
// The code does not compile: functions are not comparable.
// Practices the restrictions on a map key type.
package main_test

import "testing"

type Cache map[func(int) int]int

func (c Cache) Get(f func(int) int, name string, x int) int {
	if v, ok := c[f]; ok {
		return v
	}
	c[f] = f(x)
	return c[f]
}

func TestCache(t *testing.T) {
	calls := 0
	sq := func(x int) int { calls++; return x * x }
	c := Cache{}
	c.Get(sq, "sq", 4)
	if got := c.Get(sq, "sq", 4); got != 16 || calls != 1 {
		t.Errorf("Get = %d, calls = %d", got, calls)
	}
}
