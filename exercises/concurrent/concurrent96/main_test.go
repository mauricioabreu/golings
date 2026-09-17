// concurrent96
// Make the tests pass!

// I AM NOT DONE
//
// parallelLen computes string lengths in parallel and returns them in their
// original order via a channel of (index, result) pairs.
// Practices passing the index along with the result.
package main_test

import (
	"reflect"
	"testing"
)

type item struct {
	idx, val int
}

func parallelLen(words []string) []int {
	ch := make(chan item, len(words))
	for i, w := range words {
		go func(i int, w string) {
			ch <- item{0, len(w)}
		}(i, w)
	}
	out := make([]int, len(words))
	for range words {
		it := <-ch
		out = append(out[:0], it.val)
	}
	return out
}

func TestParallelLen(t *testing.T) {
	got := parallelLen([]string{"a", "bbb", "cc"})
	if !reflect.DeepEqual(got, []int{1, 3, 2}) {
		t.Errorf("parallelLen = %v", got)
	}
}
