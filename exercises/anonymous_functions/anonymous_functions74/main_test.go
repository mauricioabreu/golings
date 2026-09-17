// anonymous_functions74
// Make the tests pass!

// I AM NOT DONE
//
// ttlCache returns a get function caching the result of load
// for the duration ttl.
// Practices a closure over a map and a clock.
package main_test

import (
	"testing"
	"time"
)

type entry struct {
	val string
	at  time.Time
}

func ttlCache(ttl time.Duration, load func(string) string) func(key string, now time.Time) string {
	cache := map[string]entry{}
	return func(key string, now time.Time) string {
		if e, ok := cache[key]; ok {
			return e.val
		}
		v := load(key)
		cache[key] = entry{val: v}
		return v
	}
}

func TestTTLCache(t *testing.T) {
	loads := 0
	get := ttlCache(time.Minute, func(k string) string { loads++; return k + "!" })
	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	get("a", t0)
	get("a", t0.Add(30*time.Second))
	get("a", t0.Add(2*time.Minute))
	if loads != 2 {
		t.Errorf("loads = %d, want 2", loads)
	}
}
