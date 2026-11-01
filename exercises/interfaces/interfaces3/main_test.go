// interfaces3
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// asInt returns the int inside v and true, or 0 and false if v holds anything else.
// It must not panic.
func asInt(v any) (int, bool) {
	// TODO: Use a type assertion that can't panic.
	return 0, false
}

// firstString returns the first string among the values, or "" if there is none.
func firstString(values []any) string {
	// TODO: Look for the first string among the values.
	return ""
}

func TestAsInt(t *testing.T) {
	tests := []struct {
		value  any
		want   int
		wantOK bool
	}{
		{value: 42, want: 42, wantOK: true},
		{value: "42", want: 0, wantOK: false},
		{value: 4.2, want: 0, wantOK: false},
		{value: nil, want: 0, wantOK: false},
	}

	for _, tc := range tests {
		got, ok := asInt(tc.value)
		if got != tc.want || ok != tc.wantOK {
			t.Errorf("asInt(%v) = (%d, %v), want (%d, %v)", tc.value, got, ok, tc.want, tc.wantOK)
		}
	}
}

func TestFirstString(t *testing.T) {
	if got := firstString([]any{1, 2.5, "go", "gopher"}); got != "go" {
		t.Errorf("firstString = %q, want %q", got, "go")
	}
	if got := firstString([]any{1, 2.5, nil}); got != "" {
		t.Errorf("firstString without strings = %q, want %q", got, "")
	}
}
