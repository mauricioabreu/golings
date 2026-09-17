// functions88
// Make the tests pass!

// I AM NOT DONE
//
// scan hands the callback a slice that is overwritten once the callback returns.
// fields keeps those slices, so the result ends up holding damaged data.
// Practices the "valid only during the call" contract.
package main_test

import "testing"

func scan(data string, emit func([]byte)) {
	buf := make([]byte, 0, 16)
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			emit(buf)
			buf = buf[:0]
			continue
		}
		buf = append(buf, data[i])
	}
	emit(buf)
}

func fields(data string) [][]byte {
	var out [][]byte
	scan(data, func(b []byte) {
		out = append(out, b)
	})
	return out
}

func TestFields(t *testing.T) {
	got := fields("ab,cd,e")
	want := []string{"ab", "cd", "e"}
	if len(got) != len(want) {
		t.Fatalf("fields returned %d items, want %d", len(got), len(want))
	}
	for i := range want {
		if string(got[i]) != want[i] {
			t.Errorf("field %d = %q, want %q", i, got[i], want[i])
		}
	}
}
