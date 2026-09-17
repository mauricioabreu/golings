// if66
// Make the tests pass!

// I AM NOT DONE
//
// transfer logs "begin", then "lock" when a lock is needed, the work,
// and "unlock" right after it, before "end".
// A defer waits for the end of the function, not the end of the if block.
package main_test

import (
	"reflect"
	"testing"
)

func transfer(needLock bool) (log []string) {
	add := func(s string) { log = append(log, s) }
	add("begin")
	if needLock {
		add("lock")
		defer add("unlock")
		add("work")
	}
	add("end")
	return log
}

func TestTransfer(t *testing.T) {
	want := []string{"begin", "lock", "work", "unlock", "end"}
	if got := transfer(true); !reflect.DeepEqual(got, want) {
		t.Errorf("transfer(true) = %v, want %v", got, want)
	}
	if got := transfer(false); !reflect.DeepEqual(got, []string{"begin", "end"}) {
		t.Errorf("transfer(false) = %v", got)
	}
}
