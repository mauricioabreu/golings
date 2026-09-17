// functions49
// Make the tests pass!

// I AM NOT DONE
//
// applyAll applies a function to every string. Upper case is what is wanted.
// Practices passing standard library functions as values.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func applyAll(items []string, f func(string) string) []string {
	out := make([]string, len(items))
	for i, s := range items {
		out[i] = f(s)
	}
	return out
}

func upperAll(items []string) []string {
	return applyAll(items, strings.ToLower)
}

func TestUpperAll(t *testing.T) {
	if got := upperAll([]string{"go", "Rust"}); !reflect.DeepEqual(got, []string{"GO", "RUST"}) {
		t.Errorf("upperAll = %v, want [GO RUST]", got)
	}
}
