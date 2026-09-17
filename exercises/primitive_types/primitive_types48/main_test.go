// primitive_types48
// Make the tests pass!

// I AM NOT DONE
//
// luhn validates a card number with the Luhn algorithm, ignoring spaces.
// Practices digits from bytes and walking a string from right to left.
package main_test

import "testing"

func luhn(num string) bool {
	sum, count := 0, 0
	for i := len(num) - 1; i >= 0; i-- {
		c := num[i]
		if c == ' ' {
			continue
		}
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if count%2 == 0 {
			d *= 2
		}
		sum += d
		count++
	}
	return count > 1 && sum%10 == 0
}

func TestLuhn(t *testing.T) {
	cases := map[string]bool{"4539 3195 0343 6467": true, "8273 1232 7352 0569": false, "059": true, "0": false, "12a4": false}
	for in, want := range cases {
		if got := luhn(in); got != want {
			t.Errorf("luhn(%q) = %v, want %v", in, got, want)
		}
	}
}
