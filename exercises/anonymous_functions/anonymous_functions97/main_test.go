// anonymous_functions97
// Make the tests pass!

// I AM NOT DONE
//
// emailsLength picks the addresses on the example.com domain, folds them to lower
// case and adds up their lengths.
// Practices composing generic helpers with literals.
package main_test

import (
	"strings"
	"testing"
)

func mapS(s []string, f func(string) string) []string {
	out := make([]string, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}

func filterS(s []string, f func(string) bool) []string {
	var out []string
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func reduceS(s []string, init int, f func(int, string) int) int {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func emailsLength(emails []string) int {
	lower := mapS(emails, strings.ToLower)
	ours := filterS(lower, func(e string) bool { return strings.HasPrefix(e, "@example.com") })
	return reduceS(ours, 0, func(acc int, e string) int { return len(e) })
}

func TestEmailsLength(t *testing.T) {
	got := emailsLength([]string{"A@Example.com", "b@other.org", "cc@example.com"})
	if got != 27 {
		t.Errorf("emailsLength = %d, want 27", got)
	}
}
