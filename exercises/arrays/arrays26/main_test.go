// arrays26
// Make the tests pass!

// I AM NOT DONE
//
// Ring keeps the last 3 values. Push adds a value, pushing the oldest one out,
// and Items returns the values from the oldest to the newest.
// Practices a fixed array with an index taken modulo its size.
package main_test

import (
	"reflect"
	"testing"
)

type Ring struct {
	buf   [3]int
	pos   int
	count int
}

func (r *Ring) Push(v int) {
	r.buf[r.pos] = v
	r.pos++
}

func (r *Ring) Items() []int {
	out := make([]int, 0, r.count)
	start := (r.pos - r.count + len(r.buf)) % len(r.buf)
	for i := 0; i < r.count; i++ {
		out = append(out, r.buf[(start+i)%len(r.buf)])
	}
	return out
}

func TestRing(t *testing.T) {
	var r Ring
	for _, v := range []int{1, 2, 3, 4, 5} {
		r.Push(v)
	}
	if got := r.Items(); !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Errorf("Items = %v, want [3 4 5]", got)
	}
}
