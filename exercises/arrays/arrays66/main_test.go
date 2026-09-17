// arrays66
// Make the tests pass!

// I AM NOT DONE
//
// Palette is an array of colours. Its Set method must change a colour by index.
// The palette is unchanged after the call.
// A method with a value receiver works on a copy of the array.
package main_test

import "testing"

type Palette [3]string

func (p Palette) Set(i int, c string) {
	p[i] = c
}

func TestPaletteSet(t *testing.T) {
	p := Palette{"red", "green", "blue"}
	p.Set(1, "yellow")
	if p[1] != "yellow" {
		t.Errorf("p[1] = %s, want yellow", p[1])
	}
}
