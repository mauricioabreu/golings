// maps7
// Make the tests pass!

// I AM NOT DONE
//
// loadTags returns a map of tags; for empty input it returns nil,
// and the caller panics when adding a tag.
// A nil map can be read from, but not written to.
package main_test

import (
	"strings"
	"testing"
)

func loadTags(s string) map[string]bool {
	if s == "" {
		return nil
	}
	tags := map[string]bool{}
	for _, t := range strings.Split(s, ",") {
		tags[t] = true
	}
	return tags
}

func TestLoadTags(t *testing.T) {
	tags := loadTags("")
	tags["new"] = true
	if len(tags) != 1 {
		t.Errorf("tags = %v", tags)
	}
	if len(loadTags("a,b")) != 2 {
		t.Errorf("loadTags(a,b) wrong")
	}
}
