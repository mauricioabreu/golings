// primitive_types44
// Make the tests pass!

// I AM NOT DONE
//
// rle encodes a string: "aaabcc" -> "a3b1c2".
// Practices comparing neighbouring characters and strconv.Itoa.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func rle(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		b.WriteByte(s[i])
		b.WriteString(strconv.Itoa(j))
		i++
	}
	return b.String()
}

func TestRLE(t *testing.T) {
	cases := map[string]string{"aaabcc": "a3b1c2", "x": "x1", "": "", "zzzzzzzzzzzz": "z12"}
	for in, want := range cases {
		if got := rle(in); got != want {
			t.Errorf("rle(%q) = %q, want %q", in, got, want)
		}
	}
}
