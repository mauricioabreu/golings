// maps62
// Make the tests pass!

// I AM NOT DONE
//
// fromRoman turns a roman numeral into an integer with a table of values.
// Practices map[byte]int and a comparison with the next character.
package main_test

import "testing"

var roman = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 100,
}

func fromRoman(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := roman[s[i]]
		total += v
	}
	return total
}

func TestFromRoman(t *testing.T) {
	cases := map[string]int{"MCMXCIV": 1994, "LVIII": 58, "IV": 4, "MMXXIV": 2024}
	for in, want := range cases {
		if got := fromRoman(in); got != want {
			t.Errorf("fromRoman(%s) = %d, want %d", in, got, want)
		}
	}
}
