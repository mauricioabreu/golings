// maps19
// Make the tests pass!

// I AM NOT DONE
//
// cachedSquare looks in the cache first and only computes afterwards.
// Practices comma-ok when working with a cache.
package main_test

import "testing"

var computed int

func cachedSquare(cache map[int]int, n int) int {
	if v, ok := cache[n]; ok {
		_ = v
	}
	computed++
	cache[n] = n * n
	return cache[n]
}

func TestCachedSquare(t *testing.T) {
	computed = 0
	cache := map[int]int{}
	if cachedSquare(cache, 4) != 16 || cachedSquare(cache, 4) != 16 {
		t.Errorf("wrong square")
	}
	if computed != 1 {
		t.Errorf("computed %d times, want 1", computed)
	}
}
