// maps24
// Make the tests pass!

// I AM NOT DONE
//
// revoke removes a token and reports whether it was there.
// delete on a missing key is not an error.
package main_test

import "testing"

func revoke(tokens map[string]bool, tok string) bool {
	_, ok := tokens[tok]
	delete(tokens, tok)
	return true
}

func TestRevoke(t *testing.T) {
	tokens := map[string]bool{"abc": true}
	if !revoke(tokens, "abc") || revoke(tokens, "abc") || revoke(tokens, "zzz") {
		t.Errorf("revoke works incorrectly")
	}
}
