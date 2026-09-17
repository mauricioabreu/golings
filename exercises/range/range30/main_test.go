// range30
// Make the tests pass!

// I AM NOT DONE
//
// parsePairs reads arguments such as ["-k", "v", "-x", "y"]: the value after a flag
// has to be skipped. Incrementing i inside the range does not work.
// The index variable of a range is overwritten on every iteration.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func parsePairs(args []string) map[string]string {
	m := map[string]string{}
	for i, a := range args {
		if strings.HasPrefix(a, "-") && i+1 < len(args) {
			m[a[1:]] = args[i+1]
			i++
		} else {
			m[a] = ""
		}
	}
	return m
}

func TestParsePairs(t *testing.T) {
	got := parsePairs([]string{"-k", "v", "-x", "y", "file"})
	if !reflect.DeepEqual(got, map[string]string{"k": "v", "x": "y", "file": ""}) {
		t.Errorf("parsePairs = %v", got)
	}
}
