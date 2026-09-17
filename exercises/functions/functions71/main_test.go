// functions71
// Make the tests pass!

// I AM NOT DONE
//
// logf must hand its arguments to fmt.Sprint so they are printed
// as separate values rather than as one slice.
// Practices the difference between f(args) and f(args...) for ...interface{}.
package main_test

import (
	"fmt"
	"testing"
)

func logf(args ...interface{}) string {
	return fmt.Sprint(args)
}

func TestLogf(t *testing.T) {
	if got := logf("id=", 7); got != "id=7" {
		t.Errorf("logf(\"id=\", 7) = %q, want %q", got, "id=7")
	}
}
