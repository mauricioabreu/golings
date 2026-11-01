// strings2
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// countCharacters returns how many characters the text has, not how many bytes.
func countCharacters(text string) int {
	// TODO: Count characters, not bytes.
	return len(text)
}

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{text: "", want: 0},
		{text: "hello", want: 5},
		{text: "héllo", want: 5},
		{text: "世界", want: 2},
		{text: "héllo, 世界", want: 9},
	}

	for _, tc := range tests {
		if got := countCharacters(tc.text); got != tc.want {
			t.Errorf("countCharacters(%q) = %d, want %d", tc.text, got, tc.want)
		}
	}
}
