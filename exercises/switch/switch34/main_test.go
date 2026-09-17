// switch34
// Make the tests pass!

// I AM NOT DONE
//
// parseRoman turns a string of roman numerals into a number, subtraction included (IV = 4).
// Practices a switch inside a loop and a comparison with the next character.
package main_test

import "testing"

func value(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 100
	}
	return 0
}

func parseRoman(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := value(s[i])
		total += v
	}
	return total
}

func TestParseRoman(t *testing.T) {
	cases := map[string]int{"III": 3, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "LXXVIII": 78, "CXLIV": 144}
	for in, want := range cases {
		if got := parseRoman(in); got != want {
			t.Errorf("parseRoman(%s) = %d, want %d", in, got, want)
		}
	}
}
