// concurrent52
// Make the tests pass!

// I AM NOT DONE
//
// primes builds a pipeline of filters and returns the first n prime numbers.
// Practices a goroutine pipeline that grows.
package main_test

import (
	"reflect"
	"testing"
)

func generate(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case ch <- i:
			case <-done:
				return
			}
		}
	}()
	return ch
}

func filter(done <-chan struct{}, in <-chan int, p int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v%p == 0 {
				select {
				case out <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return out
}

func primes(n int) []int {
	done := make(chan struct{})
	defer close(done)
	ch := generate(done)
	var out []int
	for i := 0; i < n; i++ {
		p := <-ch
		out = append(out, p)
		filter(done, ch, p)
	}
	return out
}

func TestPrimes(t *testing.T) {
	if got := primes(6); !reflect.DeepEqual(got, []int{2, 3, 5, 7, 11, 13}) {
		t.Errorf("primes = %v", got)
	}
}
