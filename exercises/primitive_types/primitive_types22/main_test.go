// primitive_types22
// Make the tests pass!

// I AM NOT DONE
//
// status must return "enabled=true" or "enabled=false".
// Practices the %t verb for a bool.
package main_test

import (
	"fmt"
	"testing"
)

func status(on bool) string {
	return fmt.Sprintf("enabled=%d", on)
}

func TestStatus(t *testing.T) {
	if status(true) != "enabled=true" || status(false) != "enabled=false" {
		t.Errorf("status = %q, %q", status(true), status(false))
	}
}
