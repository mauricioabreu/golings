// strings3
// Make the tests pass!

package main_test

import (
	"slices"
	"testing"
)

// reverse returns the text with its characters in the opposite order.
func reverse(text string) string {
	// Reversing the bytes would tear apart every character that takes more than one.
	runes := []rune(text)
	slices.Reverse(runes)
	return string(runes)
}

func TestReverse(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{text: "", want: ""},
		{text: "hello", want: "olleh"},
		{text: "héllo", want: "olléh"},
		{text: "世界", want: "界世"},
		{text: "Go ❤ Gophers", want: "srehpoG ❤ oG"},
	}

	for _, tc := range tests {
		if got := reverse(tc.text); got != tc.want {
			t.Errorf("reverse(%q) = %q, want %q", tc.text, got, tc.want)
		}
	}
}
