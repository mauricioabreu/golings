// concurrent61
// Make the tests pass!

// I AM NOT DONE
//
// dedupe reads a stream and passes on only the first occurrence of
// each value.
// Practices a stage goroutine with its own state.
package main_test

import (
	"reflect"
	"testing"
	"time"
)

func dedupe(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		seen := map[string]bool{}
		for v := range in {
			out <- v
		}
	}()
	return out
}

func TestDedupe(t *testing.T) {
	in := make(chan string, 5)
	for _, s := range []string{"a", "b", "a", "c", "b"} {
		in <- s
	}
	close(in)
	res := make(chan []string, 1)
	go func() {
		var got []string
		for v := range dedupe(in) {
			got = append(got, v)
		}
		res <- got
	}()
	select {
	case got := <-res:
		if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
			t.Errorf("dedupe = %v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("dedupe is stuck")
	}
}
