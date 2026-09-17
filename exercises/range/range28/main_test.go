// range28
// Make the tests pass!

// I AM NOT DONE
//
// takeUntilStop has to handle the elements before "stop" and return them.
// The slice is cut inside the loop, and the loop carries on.
// A range evaluates the slice once; reassigning the variable does not change it.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func takeUntilStop(s []string) []string {
	var out []string
	for i, v := range s {
		if v == "stop" {
			s = s[:i]
			continue
		}
		_ = i
		out = append(out, strings.ToUpper(v))
	}
	return out
}

func TestTakeUntilStop(t *testing.T) {
	got := takeUntilStop([]string{"a", "b", "stop", "c"})
	if !reflect.DeepEqual(got, []string{"A", "B"}) {
		t.Errorf("takeUntilStop = %v", got)
	}
}
