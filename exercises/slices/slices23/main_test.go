// slices23
// Make the tests pass!

// I AM NOT DONE
//
// render returns the contents of a buffer that is reused afterwards.
// The data that was returned is overwritten by the next call.
// bytes.Buffer.Bytes() is only valid until the buffer changes again.
package main_test

import (
	"bytes"
	"testing"
)

type Renderer struct{ buf bytes.Buffer }

func (r *Renderer) render(name string) []byte {
	r.buf.Reset()
	r.buf.WriteString("hello, ")
	r.buf.WriteString(name)
	return r.buf.Bytes()
}

func TestRender(t *testing.T) {
	var r Renderer
	a := r.render("ann")
	b := r.render("bob")
	if string(a) != "hello, ann" || string(b) != "hello, bob" {
		t.Errorf("a = %q, b = %q", a, b)
	}
}
