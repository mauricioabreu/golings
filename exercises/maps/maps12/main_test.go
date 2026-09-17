// maps12
// Make the tests pass!

// I AM NOT DONE
//
// The price buckets are keyed by the amount as a float64. The sum 0.1+0.2
// does not find the 0.3 bucket.
// Computed float keys do not match the literals.
package main_test

import (
	"math"
	"testing"
)

type Buckets map[float64]string

func (b Buckets) Add(price float64, name string) { b[price] = name }
func (b Buckets) Get(price float64) string       { return b[price] }

func TestBuckets(t *testing.T) {
	_ = math.Round
	b := Buckets{}
	b.Add(0.3, "cheap")
	b.Add(0.29, "cheaper")
	a, c := 0.1, 0.2
	if got := b.Get(a + c); got != "cheap" {
		t.Errorf("bucket for 0.1+0.2 = %q, want cheap", got)
	}
	if got := b.Get(0.29); got != "cheaper" {
		t.Errorf("bucket for 0.29 = %q", got)
	}
}
