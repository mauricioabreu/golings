// range4
// Make the tests pass!

// I AM NOT DONE
//
// runAll calls every step function and returns their results.
// Practices a range over a slice of functions.
package main_test

import (
	"reflect"
	"testing"
)

func runAll(steps []func() string) []string {
	var out []string
	for _, step := range steps {
		out = append(out, step()[:1])
	}
	return out
}

func TestRunAll(t *testing.T) {
	steps := []func() string{
		func() string { return "build" },
		func() string { return "test" },
	}
	if got := runAll(steps); !reflect.DeepEqual(got, []string{"build", "test"}) {
		t.Errorf("runAll = %v", got)
	}
}
