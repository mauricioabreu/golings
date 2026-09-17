// functions10
// Make the tests pass!

// I AM NOT DONE
//
// hanoi(n, from, to, via) must return the list of moves for n disks.
// Practices recursion with several recursive calls.
package main_test

import (
	"reflect"
	"testing"
)

func hanoi(n int, from, to, via string) []string {
	if n == 0 {
		return nil
	}
	moves := hanoi(n-1, from, to, via)
	moves = append(moves, from+"->"+to)
	return append(moves, hanoi(n-1, from, to, via)...)
}

func TestHanoi(t *testing.T) {
	want := []string{"A->B", "A->C", "B->C"}
	if got := hanoi(2, "A", "C", "B"); !reflect.DeepEqual(got, want) {
		t.Errorf("hanoi(2) = %v, want %v", got, want)
	}
	if got := len(hanoi(10, "A", "C", "B")); got != 1023 {
		t.Errorf("hanoi(10) has %d moves, want 1023", got)
	}
}
