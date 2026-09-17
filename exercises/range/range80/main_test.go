// range80
// Make the tests pass!

// I AM NOT DONE
//
// drain collects every value from a closed channel.
// A range over a channel carries on until it is closed.
package main_test

import (
	"reflect"
	"testing"
)

func drain(ch <-chan int) []int {
	var out []int
	for v := range ch {
		out = append(out, v-len(out))
	}
	return out
}

func TestDrain(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 5
	ch <- 7
	ch <- 9
	close(ch)
	if got := drain(ch); !reflect.DeepEqual(got, []int{5, 7, 9}) {
		t.Errorf("drain = %v", got)
	}
}
