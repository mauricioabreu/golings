// variables57
// Make the tests pass!

// I AM NOT DONE
//
// csvLine must join the fields with commas.
// The zero value of a strings.Builder is ready to use without initialization.
// Practices variables with a useful zero value.
package main_test

import (
	"strings"
	"testing"
)

func csvLine(fields []string) string {
	var b strings.Builder
	for _, f := range fields {
		b.WriteString(f)
		b.WriteString(",")
	}
	return b.String()
}

func TestCSVLine(t *testing.T) {
	cases := map[string][]string{"a,b,c": {"a", "b", "c"}, "x": {"x"}, "": nil}
	for want, in := range cases {
		if got := csvLine(in); got != want {
			t.Errorf("csvLine(%v) = %q, want %q", in, got, want)
		}
	}
}
