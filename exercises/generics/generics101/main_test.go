// generics101
// Make the tests pass!

// I AM NOT DONE
//
// Mean returns the average of a slice as a float64, and false for an empty slice.
// Practices converting a T to float64.
package main_test

import "testing"

type Number interface {
	~int | ~int8 | ~int64 | ~float64
}

func Mean[T Number](s []T) (float64, bool) {
	var sum T
	for _, v := range s {
		sum += v
	}
	return float64(sum) / float64(len(s)), true
}

func TestMean(t *testing.T) {
	if m, ok := Mean([]int8{100, 100, 100}); !ok || m != 100 {
		t.Errorf("Mean(int8) = %v, %v", m, ok)
	}
	if _, ok := Mean([]float64{}); ok {
		t.Errorf("Mean(empty) should fail")
	}
}
