// anonymous_functions61
// Make the tests pass!

// I AM NOT DONE
//
// run carries out an action, calling the before and after hooks when they are given,
// and the after hook runs even when the action fails.
// Practices optional literals and defer.
package main_test

import (
	"errors"
	"reflect"
	"testing"
)

type Hooks struct {
	Before, After func()
}

func run(h Hooks, action func() error) error {
	h.Before()
	err := action()
	h.After()
	return err
}

func TestRun(t *testing.T) {
	var log []string
	h := Hooks{After: func() { log = append(log, "after") }}
	err := run(h, func() error { log = append(log, "act"); return errors.New("x") })
	if err == nil || !reflect.DeepEqual(log, []string{"act", "after"}) {
		t.Errorf("err=%v log=%v", err, log)
	}
}
