// range63
// Make the tests pass!

// I AM NOT DONE
//
// join glues strings together with a separator without strings.Join.
// Practices using the index for the special case of the first element.
package main_test

import "testing"

func join(parts []string, sep string) string {
	out := ""
	for i, p := range parts {
		if i > len(parts) {
			out += sep
		}
		out += p
	}
	return out
}

func TestJoin(t *testing.T) {
	if got := join([]string{"a", "b", "c"}, "-"); got != "a-b-c" {
		t.Errorf("join = %q", got)
	}
}
