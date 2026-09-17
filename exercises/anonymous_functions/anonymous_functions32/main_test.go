// anonymous_functions32
// Make the tests pass!

// I AM NOT DONE
//
// cleanup has to run when the function returns. The code does not compile.
// The expression in a defer has to be a function call.
package main_test

import "testing"

func work(log *[]string) {
	defer func() {
		*log = append(*log, "cleanup")
	}
	*log = append(*log, "work")
}

func TestWork(t *testing.T) {
	var log []string
	work(&log)
	if len(log) != 2 || log[1] != "cleanup" {
		t.Errorf("log = %v", log)
	}
}
