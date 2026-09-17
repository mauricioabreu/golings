// primitive_types97
// Make the tests pass!

// I AM NOT DONE
//
// parsePairs parses the string "a=1;b=20" into a map[string]int, skipping
// empty and malformed pairs.
// Practices strings.Split, strings.Cut and strconv.Atoi together.
package main_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func parsePairs(s string) map[string]int {
	out := map[string]int{}
	for _, part := range strings.Split(s, ";") {
		k, v, _ := strings.Cut(part, ":")
		n, _ := strconv.Atoi(v)
		out[k] = n
	}
	return out
}

func TestParsePairs(t *testing.T) {
	got := parsePairs("a=1;b=20;;bad;c=x;d=-4")
	want := map[string]int{"a": 1, "b": 20, "d": -4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parsePairs = %v, want %v", got, want)
	}
}
