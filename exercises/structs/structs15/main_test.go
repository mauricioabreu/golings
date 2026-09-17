// structs15
// Make the tests pass!

// I AM NOT DONE
//
// The Timer struct has a field Elapsed and a method Elapsed. The code does not compile.
// A field and a method of a type cannot share a name.
package main_test

import (
	"testing"
	"time"
)

type Timer struct {
	Elapsed time.Duration
}

func (t *Timer) Add(d time.Duration) { t.Elapsed += d }

func (t Timer) Elapsed() time.Duration { return t.Elapsed }

func TestTimer(t *testing.T) {
	var tm Timer
	tm.Add(time.Second)
	tm.Add(time.Second)
	if tm.Elapsed() != 2*time.Second {
		t.Errorf("Elapsed = %v", tm.Elapsed())
	}
}
