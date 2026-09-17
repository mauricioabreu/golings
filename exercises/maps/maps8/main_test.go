// maps8
// Make the tests pass!

// I AM NOT DONE
//
// memo wraps a function so that repeated calls with the same argument
// take the result from a cache.
// Practices a map inside a closure.
package main_test

import "testing"

func memo(f func(int) int) func(int) int {
	return func(n int) int {
		cache := map[int]int{}
		cache[n] = f(n)
		return cache[n]
	}
}

func TestMemo(t *testing.T) {
	calls := 0
	sq := memo(func(n int) int { calls++; return n * n })
	if sq(3) != 9 || sq(3) != 9 || sq(4) != 16 {
		t.Errorf("wrong results")
	}
	if calls != 2 {
		t.Errorf("underlying function called %d times, want 2", calls)
	}
}
