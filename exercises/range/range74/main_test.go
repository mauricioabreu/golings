// range74
// Make the tests pass!

// I AM NOT DONE
//
// trimAll strips the spaces around every string of a slice.
// The strings do not change: the result of the function is thrown away.
// Strings are immutable, and the strings functions return new ones.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func trimAll(lines []string) {
	for _, l := range lines {
		l = strings.TrimSpace(l)
		_ = l
	}
}

func TestTrimAll(t *testing.T) {
	lines := []string{"  a ", "b\n", "c"}
	trimAll(lines)
	if !reflect.DeepEqual(lines, []string{"a", "b", "c"}) {
		t.Errorf("trimAll = %q", lines)
	}
}
