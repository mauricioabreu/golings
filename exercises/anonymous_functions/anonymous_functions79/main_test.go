// anonymous_functions79
// Make the tests pass!

// I AM NOT DONE
//
// report starts a goroutine that has to send the original name.
// The goroutine reads the variable later, once it has changed.
// A literal in a goroutine reads the variable at the moment it runs.
package main_test

import "testing"

func report() string {
	name := "draft"
	ready := make(chan struct{})
	out := make(chan string)
	go func() {
		<-ready
		out <- name
	}()
	name = "final"
	close(ready)
	return <-out
}

func TestReport(t *testing.T) {
	if got := report(); got != "draft" {
		t.Errorf("report = %q, want draft", got)
	}
}
