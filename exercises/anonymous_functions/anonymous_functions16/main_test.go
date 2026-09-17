// anonymous_functions16
// Make the tests pass!

// I AM NOT DONE
//
// getOr returns a value, or calls a literal computing the default one.
// Right now the literal is called every time, even when the value is there.
// Practices passing a function for a lazy computation.
package main_test

import "testing"

func getOr(m map[string]int, k string, def func() int) int {
	fallback := def()
	if v, ok := m[k]; ok {
		return v
	}
	return fallback
}

func TestGetOr(t *testing.T) {
	calls := 0
	def := func() int { calls++; return 42 }
	m := map[string]int{"a": 1}
	if getOr(m, "a", def) != 1 || getOr(m, "b", def) != 42 {
		t.Errorf("getOr works incorrectly")
	}
	if calls != 1 {
		t.Errorf("default computed %d times", calls)
	}
}
