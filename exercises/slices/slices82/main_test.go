// slices82
// Make the tests pass!

// I AM NOT DONE
//
// histogram counts the words and returns (word, count) pairs sorted
// by count in descending order, and by word on a tie.
// Practices a slice of structs built from a map plus sort.Slice.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type WordCount struct {
	Word  string
	Count int
}

func histogram(words []string) []WordCount {
	counts := map[string]int{}
	for _, w := range words {
		counts[w]++
	}
	var out []WordCount
	for w, c := range counts {
		out = append(out, WordCount{w, c})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Count < out[j].Count
	})
	return out
}

func TestHistogram(t *testing.T) {
	got := histogram([]string{"b", "a", "c", "a", "b", "a"})
	want := []WordCount{{"a", 3}, {"b", 2}, {"c", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("histogram = %v", got)
	}
	got = histogram([]string{"y", "x"})
	if !reflect.DeepEqual(got, []WordCount{{"x", 1}, {"y", 1}}) {
		t.Errorf("histogram ties = %v", got)
	}
}
