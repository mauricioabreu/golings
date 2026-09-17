// generics82
// Make the tests pass!

// I AM NOT DONE
//
// ParseAll builds values of type T and calls their Set method, which has a pointer receiver.
// The code does not compile: T has no Set method, only *T does.
// Practices the "pointer to T with a method" constraint pattern.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

type Tag struct{ Name string }

func (t *Tag) Set(s string) { t.Name = strings.ToLower(s) }

type Setter interface{ Set(string) }

func ParseAll[T Setter](items []string) []T {
	out := make([]T, len(items))
	for i, s := range items {
		out[i].Set(s)
	}
	return out
}

func tags() []Tag { return ParseAll[Tag]([]string{"Go", "WEB"}) }

func TestTags(t *testing.T) {
	if got := tags(); !reflect.DeepEqual(got, []Tag{{"go"}, {"web"}}) {
		t.Errorf("tags = %v", got)
	}
}
