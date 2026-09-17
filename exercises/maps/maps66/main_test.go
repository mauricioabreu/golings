// maps66
// Make the tests pass!

// I AM NOT DONE
//
// sortedPairs returns "key=value" pairs in key order.
// The keys come out sorted, but the values do not match them.
// The iteration order of a map has to be taken into account for keys and values together.
package main_test

import (
	"sort"
	"strconv"
	"testing"
)

func sortedPairs(m map[string]int) []string {
	var keys []string
	var vals []int
	for k, v := range m {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	sort.Strings(keys)
	out := make([]string, len(keys))
	for i, k := range keys {
		out[i] = k + "=" + strconv.Itoa(vals[i])
	}
	return out
}

func TestSortedPairs(t *testing.T) {
	m := map[string]int{"e": 5, "a": 1, "d": 4, "b": 2, "c": 3, "f": 6, "g": 7, "h": 8}
	for i := 0; i < 30; i++ {
		got := sortedPairs(m)
		for j, p := range got {
			want := string(rune('a'+j)) + "=" + strconv.Itoa(j+1)
			if p != want {
				t.Fatalf("pair %d = %q, want %q (all: %v)", j, p, want, got)
			}
		}
	}
}
