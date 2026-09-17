// concurrent55
// Make the tests pass!

// I AM NOT DONE
//
// The reading function expects a <-chan but is given a send-only channel.
// The code does not compile.
// Practices that chan<- cannot be turned into <-chan.
package main_test

import "testing"

func fill(out chan<- string) {
	out <- "a"
	out <- "b"
	close(out)
}

func count(in <-chan string) int {
	n := 0
	for range in {
		n++
	}
	return n
}

func pipeline() int {
	ch := make(chan string, 2)
	var send chan<- string = ch
	fill(send)
	return count(send)
}

func TestPipeline(t *testing.T) {
	if pipeline() != 2 {
		t.Errorf("pipeline = %d", pipeline())
	}
}
