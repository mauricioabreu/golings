// testing2
// Make the tests pass!

package main_test

import (
	"fmt"
	"testing"
)

// assertEqual fails the test when got and want differ. The failure should be
// reported at the line of the test that called assertEqual, not at a line
// inside this helper.
func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// fakeTB stands in for the testing.T the helper receives, to see what the
// helper does with it.
type fakeTB struct {
	testing.TB
	helper bool
	failed []string
	logged []string
}

func (f *fakeTB) Helper() { f.helper = true }

func (f *fakeTB) Logf(format string, args ...any) {
	f.logged = append(f.logged, fmt.Sprintf(format, args...))
}

func (f *fakeTB) Errorf(format string, args ...any) {
	f.failed = append(f.failed, fmt.Sprintf(format, args...))
}

func (f *fakeTB) Fatalf(format string, args ...any) {
	f.failed = append(f.failed, fmt.Sprintf(format, args...))
}

func TestAssertEqualFailsTheTest(t *testing.T) {
	fake := &fakeTB{}

	assertEqual(fake, 1, 2)

	if len(fake.failed) != 1 {
		t.Errorf("assertEqual(1, 2) failed the test %d times, want once", len(fake.failed))
	}
}

func TestAssertEqualIsATestHelper(t *testing.T) {
	fake := &fakeTB{}

	assertEqual(fake, 1, 2)

	if !fake.helper {
		t.Error("assertEqual should call t.Helper()")
	}
}

func TestAssertEqualPassesWhenEqual(t *testing.T) {
	fake := &fakeTB{}

	assertEqual(fake, 2, 2)

	if len(fake.failed) != 0 {
		t.Errorf("assertEqual(2, 2) failed the test: %v", fake.failed)
	}
}
