// concurrent86
// Make the tests pass!

// I AM NOT DONE
//
// produce is declared with a send-only channel but tries to read from it.
// The code does not compile.
// Practices directional channels chan<- and <-chan.
package main_test

import "testing"

func produce(out chan<- int) {
	for i := 0; i < 3; i++ {
		out <- i
	}
	<-out
	close(out)
}

func TestProduce(t *testing.T) {
	ch := make(chan int, 3)
	produce(ch)
	sum := 0
	for v := range ch {
		sum += v
	}
	if sum != 3 {
		t.Errorf("sum = %d", sum)
	}
}
