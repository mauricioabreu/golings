// interfaces4
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// describe says what kind of value it got:
//
//	an int 42          -> "an int: 42"
//	a string "hi"      -> "a string: hi"
//	a bool true        -> "a bool: true"
//	a slice of ints    -> "a slice of 3 ints"
//	nil                -> "nothing"
//	anything else      -> "something else"
func describe(v any) string {
	return "something else"
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
