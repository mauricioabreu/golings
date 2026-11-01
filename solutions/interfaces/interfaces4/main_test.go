// interfaces4
// Make the tests pass!

package main_test

import (
	"fmt"
	"testing"
)

// describe says what kind of value it got:
//
//	an int 42          -> "an int: 42"
//	a string "hi"      -> "a string: hi"
//	a bool true        -> "a bool: true"
//	a slice of ints    -> "a slice of 3 ints"
//	nil                -> "nothing"
//	anything else      -> "something else"
func describe(v any) string {
	switch v := v.(type) {
	case nil:
		return "nothing"
	case int:
		return fmt.Sprintf("an int: %d", v)
	case string:
		return "a string: " + v
	case bool:
		return fmt.Sprintf("a bool: %t", v)
	case []int:
		return fmt.Sprintf("a slice of %d ints", len(v))
	default:
		return "something else"
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		value any
		want  string
	}{
		{value: 42, want: "an int: 42"},
		{value: "hi", want: "a string: hi"},
		{value: true, want: "a bool: true"},
		{value: []int{1, 2, 3}, want: "a slice of 3 ints"},
		{value: nil, want: "nothing"},
		{value: 4.2, want: "something else"},
	}

	for _, tc := range tests {
		if got := describe(tc.value); got != tc.want {
			t.Errorf("describe(%v) = %q, want %q", tc.value, got, tc.want)
		}
	}
}
