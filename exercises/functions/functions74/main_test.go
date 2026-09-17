// functions74
// Make the tests pass!

// I AM NOT DONE
//
// joinAll must pass every string to fmt.Sprint as a separate argument.
// The code does not compile.
// A []T cannot be passed as ...interface{}, even when T would fit.
package main_test

import (
	"fmt"
	"testing"
)

func joinAll(names []string) string {
	return fmt.Sprint(names...)
}

func TestJoinAll(t *testing.T) {
	if got := joinAll([]string{"go", "-", "pher"}); got != "go-pher" {
		t.Errorf("joinAll = %q, want %q", got, "go-pher")
	}
}
