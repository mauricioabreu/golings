// generics31
// Make the tests pass!

// I AM NOT DONE
//
// UniqBy keeps the first element for every key value.
// Practices a generic T with a comparable K.
package main_test

import (
	"reflect"
	"testing"
)

func UniqBy[T any, K comparable](s []T, key func(T) K) []T {
	seen := map[K]bool{}
	var out []T
	for _, v := range s {
		out = append(out, v)
		seen[key(v)] = true
	}
	return out
}

type Mail struct{ From, Subject string }

func TestUniqBy(t *testing.T) {
	ms := []Mail{{"ann", "hi"}, {"bob", "yo"}, {"ann", "again"}}
	got := UniqBy(ms, func(m Mail) string { return m.From })
	if !reflect.DeepEqual(got, []Mail{{"ann", "hi"}, {"bob", "yo"}}) {
		t.Errorf("UniqBy = %v", got)
	}
}
