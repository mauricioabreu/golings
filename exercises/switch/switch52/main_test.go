// switch52
// Make the tests pass!

// I AM NOT DONE
//
// tokenKind works out the kind of a token: "keyword" (if, for, func), "number" (digits only),
// "ident" (anything else that is not empty) or "empty".
// Practices a switch with function calls in the conditions.
package main_test

import (
	"strings"
	"testing"
)

func isNumber(s string) bool {
	return s != "" && strings.Trim(s, "0123456789") == ""
}

func tokenKind(tok string) string {
	switch {
	case tok == "if" || tok == "for":
		return "keyword"
	case !isNumber(tok):
		return "number"
	}
	return "ident"
}

func TestTokenKind(t *testing.T) {
	cases := map[string]string{"if": "keyword", "func": "keyword", "42": "number", "x1": "ident", "": "empty"}
	for in, want := range cases {
		if got := tokenKind(in); got != want {
			t.Errorf("tokenKind(%q) = %s, want %s", in, got, want)
		}
	}
}
