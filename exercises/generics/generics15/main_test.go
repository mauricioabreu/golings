// generics15
// Make the tests pass!

// I AM NOT DONE
//
// Memoize wraps a function, caching the results by argument.
// Practices generic closures.
package main_test

import (
	"strings"
	"testing"
)

func Memoize[K comparable, V any](f func(K) V) func(K) V {
	cache := map[K]V{}
	return func(k K) V {
		cache[k] = f(k)
		return cache[k]
	}
}

func TestMemoize(t *testing.T) {
	calls := 0
	up := Memoize(func(s string) string { calls++; return strings.ToUpper(s) })
	if up("go") != "GO" || up("go") != "GO" || up("c") != "C" || calls != 2 {
		t.Errorf("calls = %d", calls)
	}
}
