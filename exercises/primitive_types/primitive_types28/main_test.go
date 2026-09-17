// primitive_types28
// Make the tests pass!

// I AM NOT DONE
//
// tags parses a comma separated string of tags. An empty string must give
// no tags at all, yet it gives one empty tag.
// strings.Split("", ",") returns [""], not an empty slice.
package main_test

import (
	"strings"
	"testing"
)

func tags(s string) []string {
	return strings.Split(s, ",")
}

func TestTags(t *testing.T) {
	if got := tags(""); len(got) != 0 {
		t.Errorf("tags(\"\") = %q, want empty", got)
	}
	if got := tags("go,web"); len(got) != 2 || got[0] != "go" || got[1] != "web" {
		t.Errorf("tags(go,web) = %q", got)
	}
	if got := tags("go,,web,"); len(got) != 2 {
		t.Errorf("tags(go,,web,) = %q, want 2 tags", got)
	}
}
