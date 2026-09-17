// primitive_types102
// Make the tests pass!

// I AM NOT DONE
//
// quoted must return the string in double quotes with the escapes applied.
// Practices the %q verb.
package main_test

import (
	"fmt"
	"testing"
)

func quoted(s string) string {
	return fmt.Sprintf("%s", s)
}

func TestQuoted(t *testing.T) {
	if got := quoted("say \"hi\""); got != `"say \"hi\""` {
		t.Errorf("quoted = %s", got)
	}
}
