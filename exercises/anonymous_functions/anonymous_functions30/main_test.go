// anonymous_functions30
// Make the tests pass!

// I AM NOT DONE
//
// newQueue returns a push and a pop function working on a shared slice.
// Practices a pair of closures over one slice.
package main_test

import "testing"

func newQueue() (push func(int), pop func() (int, bool)) {
	var items []int
	push = func(v int) { items = append(items, v) }
	pop = func() (int, bool) {
		v := items[len(items)-1]
		return v, true
	}
	return
}

func TestQueue(t *testing.T) {
	push, pop := newQueue()
	push(1)
	push(2)
	a, _ := pop()
	b, _ := pop()
	_, ok := pop()
	if a != 1 || b != 2 || ok {
		t.Errorf("pop = %d %d %v", a, b, ok)
	}
}
