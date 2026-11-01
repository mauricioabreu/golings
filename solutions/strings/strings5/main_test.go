// strings5
// Make the tests pass!

package main_test

import (
	"strings"
	"testing"
)

// wordCount returns how many words the text has. Words are separated by any
// amount of whitespace: spaces, tabs, new lines.
func wordCount(text string) int {
	return len(strings.Fields(text))
}

// normalize returns the name without the spaces around it, in lower case.
func normalize(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{text: "", want: 0},
		{text: "hello world", want: 2},
		{text: "  spaced   out  ", want: 2},
		{text: "line one\nline two", want: 4},
		{text: "tab\tseparated", want: 2},
	}

	for _, tc := range tests {
		if got := wordCount(tc.text); got != tc.want {
			t.Errorf("wordCount(%q) = %d, want %d", tc.text, got, tc.want)
		}
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "gopher", want: "gopher"},
		{name: "  Gopher ", want: "gopher"},
		{name: "\tADA Lovelace\n", want: "ada lovelace"},
	}

	for _, tc := range tests {
		if got := normalize(tc.name); got != tc.want {
			t.Errorf("normalize(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}
