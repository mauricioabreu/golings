// strings4
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// repeatWithSeparator returns the word n times, with sep between the copies:
// ("go", 3, "-") gives "go-go-go".
func repeatWithSeparator(word string, n int, sep string) string {
	// TODO: Build the result without copying the whole string on every step.
	result := ""
	for i := 0; i < n; i++ {
		if i > 0 {
			result += sep
		}
		result += word
	}
	return result
}

func TestRepeatWithSeparator(t *testing.T) {
	tests := []struct {
		word string
		n    int
		sep  string
		want string
	}{
		{word: "go", n: 3, sep: "-", want: "go-go-go"},
		{word: "a", n: 1, sep: ",", want: "a"},
		{word: "a", n: 0, sep: ",", want: ""},
		{word: "ab", n: 2, sep: "", want: "abab"},
	}

	for _, tc := range tests {
		if got := repeatWithSeparator(tc.word, tc.n, tc.sep); got != tc.want {
			t.Errorf("repeatWithSeparator(%q, %d, %q) = %q, want %q", tc.word, tc.n, tc.sep, got, tc.want)
		}
	}
}

// The result is right, but every += copies the whole string built so far into a
// new one. Building a big string that way is slow.
func TestRepeatWithSeparatorDoesNotCopyOnEveryStep(t *testing.T) {
	allocations := testing.AllocsPerRun(10, func() {
		repeatWithSeparator("go", 200, "-")
	})

	if allocations > 20 {
		t.Errorf("repeatWithSeparator allocated %.0f times for 200 words, want 20 or fewer: use a strings.Builder", allocations)
	}
}
