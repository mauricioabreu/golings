// functions41
// Make the tests pass!

// I AM NOT DONE
//
// handler returns a handler function; the cleanup must run
// after the handling, on every call of the handler.
// A defer belongs to the function whose body it is written in.
package main_test

import (
	"reflect"
	"testing"
)

func handler(log *[]string) func() {
	*log = append(*log, "setup")
	defer func() { *log = append(*log, "teardown") }()
	return func() {
		*log = append(*log, "handle")
	}
}

func TestHandler(t *testing.T) {
	var log []string
	h := handler(&log)
	h()
	h()
	want := []string{"setup", "handle", "teardown", "handle", "teardown"}
	if !reflect.DeepEqual(log, want) {
		t.Errorf("log = %v, want %v", log, want)
	}
}
