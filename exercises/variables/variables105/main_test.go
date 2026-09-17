// variables105
// Make the tests pass!

// I AM NOT DONE
//
// Width=0, Height=1 and Depth=2 are expected.
// Right now two constants share one specification line.
// iota grows per line (per ConstSpec), not per name.
package main_test

import "testing"

const (
	Width, Height = iota, iota
	Depth         = iota
)

func TestAxes(t *testing.T) {
	if Width != 0 || Height != 1 || Depth != 2 {
		t.Errorf("Width=%d Height=%d Depth=%d, want 0 1 2", Width, Height, Depth)
	}
}
