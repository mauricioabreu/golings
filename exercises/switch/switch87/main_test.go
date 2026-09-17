// switch87
// Make the tests pass!

// I AM NOT DONE
//
// describe turns a score of 1-5 into a word and returns an error for anything else.
// Practices a switch returning an error from its default.
package main_test

import (
	"errors"
	"testing"
)

func describe(score int) (string, error) {
	switch score {
	case 1, 2:
		return "bad", nil
	case 3:
		return "ok", nil
	case 4:
		return "great", nil
	}
	return "", nil
}

func TestDescribe(t *testing.T) {
	_ = errors.New
	cases := map[int]string{1: "bad", 2: "bad", 3: "ok", 4: "great", 5: "great"}
	for in, want := range cases {
		if got, err := describe(in); err != nil || got != want {
			t.Errorf("describe(%d) = %s, %v", in, got, err)
		}
	}
	if _, err := describe(0); err == nil {
		t.Errorf("describe(0) should fail")
	}
}
