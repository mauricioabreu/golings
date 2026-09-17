// generics21
// Make the tests pass!

// I AM NOT DONE
//
// Avg computes the average, converting the sum to a float64.
// Practices converting the values of a type parameter.
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Avg[T Number](s []T) float64 {
	var sum T
	for _, v := range s {
		sum += v
	}
	return float64(sum) / 2
}

func TestAvg(t *testing.T) {
	if Avg([]int{1, 2, 3, 4, 5}) != 3 {
		t.Errorf("Avg = %v", Avg([]int{1, 2, 3, 4, 5}))
	}
}
