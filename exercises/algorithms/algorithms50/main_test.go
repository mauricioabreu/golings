// algorithms50
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: backtracking with constraints. Return all valid bracket
// sequences of n pairs (in lexicographic order when ')' follows '(').
// Expected asymptotics: O(4^n / √n) time, O(n) space.
package main_test

import (
	"reflect"
	"testing"
)

func generateParenthesis(n int) []string {
	return nil
}

func TestGenerateParenthesis(t *testing.T) {
	if got := generateParenthesis(3); !reflect.DeepEqual(got, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}) {
		t.Errorf("generateParenthesis(3) = %v", got)
	}
	if got := generateParenthesis(1); !reflect.DeepEqual(got, []string{"()"}) {
		t.Errorf("generateParenthesis(1) = %v", got)
	}
	if got := generateParenthesis(0); len(got) != 1 || got[0] != "" {
		t.Errorf("generateParenthesis(0) = %v", got)
	}
}
