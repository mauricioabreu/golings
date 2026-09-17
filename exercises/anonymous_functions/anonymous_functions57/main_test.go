// anonymous_functions57
// Make the tests pass!

// I AM NOT DONE
//
// save has to roll the changes back when a step fails.
// The deferred literal checks err, and the error went into a shadowed variable.
// A literal sees the outer variable, not an inner one of the same name.
package main_test

import (
	"errors"
	"testing"
)

func save(fail bool, log *[]string) error {
	var err error
	defer func() {
		if err != nil {
			*log = append(*log, "rollback")
		}
	}()
	if fail {
		err := errors.New("disk full")
		return err
	}
	*log = append(*log, "commit")
	return nil
}

func TestSave(t *testing.T) {
	var log []string
	if save(true, &log) == nil || len(log) != 1 || log[0] != "rollback" {
		t.Errorf("failed save log = %v", log)
	}
}
