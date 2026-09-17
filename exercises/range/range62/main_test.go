// range62
// Make the tests pass!

// I AM NOT DONE
//
// firstN reads at most n values from a channel and reports whether the channel
// was closed early.
// Practices v, ok := <-ch in a loop as an alternative to range.
package main_test

import (
	"reflect"
	"testing"
)

func firstN(ch <-chan int, n int) ([]int, bool) {
	var out []int
	for v := range ch {
		out = append(out, v)
	}
	return out, true
}

func TestFirstN(t *testing.T) {
	ch := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	got, closed := firstN(ch, 3)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) || closed {
		t.Errorf("firstN = %v, %v", got, closed)
	}
	close(ch)
	got, closed = firstN(ch, 5)
	if !reflect.DeepEqual(got, []int{4, 5}) || !closed {
		t.Errorf("firstN after close = %v, %v", got, closed)
	}
}
