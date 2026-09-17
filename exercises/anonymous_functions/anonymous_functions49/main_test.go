// anonymous_functions49
// Make the tests pass!

// I AM NOT DONE
//
// Editor holds the text and a stack of undo literals; Undo rolls the last action back.
// Practices storing inverse operations as closures.
package main_test

import "testing"

type Editor struct {
	text string
	undo []func()
}

func (e *Editor) Append(s string) {
	prev := e.text
	e.undo = append(e.undo, func() { e.text = "" })
	_ = prev
	e.text += s
}

func (e *Editor) Undo() {
	if len(e.undo) == 0 {
		return
	}
	e.undo[0]()
	e.undo = e.undo[1:]
}

func TestEditor(t *testing.T) {
	var e Editor
	e.Append("go")
	e.Append(" is")
	e.Append(" fun")
	e.Undo()
	if e.text != "go is" {
		t.Errorf("after one undo text = %q", e.text)
	}
	e.Undo()
	e.Undo()
	e.Undo()
	if e.text != "" {
		t.Errorf("after all undos text = %q", e.text)
	}
}
