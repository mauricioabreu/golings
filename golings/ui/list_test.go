package ui_test

import (
	"bytes"
	"testing"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	. "github.com/onsi/gomega"
)

func TestPrintList(t *testing.T) {
	g := NewWithT(t)
	var out bytes.Buffer

	ui.PrintList(&out, []exercises.Exercise{
		{Name: "pending1", Path: "../fixtures/pending1/main.go"},
		{Name: "success1", Path: "../fixtures/success1/main.go"},
	})

	g.Expect(out.String()).To(And(
		ContainSubstring("NAME"),
		ContainSubstring("STATE"),
		MatchRegexp(`pending1 \| \S+\s+\| +Pending`),
		MatchRegexp(`success1 \| \S+\s+\| +Done`),
	))
}

func TestPrintListEmpty(t *testing.T) {
	g := NewWithT(t)
	var out bytes.Buffer

	ui.PrintList(&out, nil)

	g.Expect(out.String()).To(ContainSubstring("NAME"))
}
