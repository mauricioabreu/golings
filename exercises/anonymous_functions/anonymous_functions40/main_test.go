// anonymous_functions40
// Make the tests pass!

// I AM NOT DONE
//
// stripVowels has to remove the vowels. The literal returns 0, and the string
// ends up holding NUL characters.
// strings.Map drops a character only on a negative result.
package main_test

import (
	"strings"
	"testing"
)

func stripVowels(s string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune("aeiou", r) {
			return 0
		}
		return r
	}, s)
}

func TestStripVowels(t *testing.T) {
	if got := stripVowels("gopher"); got != "gphr" {
		t.Errorf("stripVowels = %q", got)
	}
}
