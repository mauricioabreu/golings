// algorithms43
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: parsing by digit groups. Write a non-negative number in English
// words, as is customary: groups of three digits with the suffixes
// Thousand, Million, Billion.
// Expected asymptotics: O(log n) time, O(1) space.
package main_test

import (
	"strings"
	"testing"
)

func numberToWords(n int) string {
	return ""
}

func TestNumberToWords(t *testing.T) {
	_ = strings.Join
	cases := map[int]string{
		0:          "Zero",
		123:        "One Hundred Twenty Three",
		12345:      "Twelve Thousand Three Hundred Forty Five",
		1234567:    "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		1000000:    "One Million",
		20:         "Twenty",
		1000010:    "One Million Ten",
		2147483647: "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven",
	}
	for in, want := range cases {
		if got := numberToWords(in); got != want {
			t.Errorf("numberToWords(%d) = %q, want %q", in, got, want)
		}
	}
}
