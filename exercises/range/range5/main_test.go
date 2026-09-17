// range5
// Make the tests pass!

// I AM NOT DONE
//
// readings sends pointers to one and the same struct into a channel, changing it as it goes.
// The receiver keeps the pointers, and every stored value comes out the same.
// A range over a channel of pointers yields references to a shared object.
package main_test

import "testing"

type Reading struct{ Value int }

func readings(n int) <-chan *Reading {
	ch := make(chan *Reading, n)
	r := &Reading{}
	for i := 0; i < n; i++ {
		r.Value = i * 10
		ch <- r
	}
	close(ch)
	return ch
}

func TestReadings(t *testing.T) {
	var got []*Reading
	for r := range readings(3) {
		got = append(got, r)
	}
	if len(got) != 3 || got[0].Value != 0 || got[1].Value != 10 || got[2].Value != 20 {
		t.Errorf("values = %d %d %d", got[0].Value, got[1].Value, got[2].Value)
	}
}
