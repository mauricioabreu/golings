// range67
// Make the tests pass!

// I AM NOT DONE
//
// countPending walks the tasks of a store. A plain for calls store.Tasks()
// on every iteration, and the request counter grows.
// A range expression is evaluated once, while a for condition is evaluated every time.
package main_test

import "testing"

type Store struct{ calls int }

func (s *Store) Tasks() []string {
	s.calls++
	return []string{"done", "todo", "todo"}
}

func countPending(s *Store) int {
	n := 0
	for i := 0; i < len(s.Tasks()); i++ {
		if s.Tasks()[i] == "todo" {
			n++
		}
	}
	return n
}

func TestCountPending(t *testing.T) {
	s := &Store{}
	if got := countPending(s); got != 2 {
		t.Errorf("countPending = %d, want 2", got)
	}
	if s.calls != 1 {
		t.Errorf("Tasks() called %d times, want 1", s.calls)
	}
}
