package ui_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	. "github.com/onsi/gomega"
)

func TestPrintTopicProgress(t *testing.T) {
	g := NewWithT(t)
	var out bytes.Buffer

	ui.PrintTopicProgress(&out, []exercises.TopicProgress{
		{Topic: "variables", Done: 3, Total: 6},
		{Topic: "slices", Done: 4, Total: 4},
		{Topic: "maps", Done: 0, Total: 3},
		{Topic: "", Done: 1, Total: 2},
	})

	g.Expect(out.String()).To(And(
		MatchRegexp(`variables\s+\| 3/6\s+\| `+strings.Repeat("█", 10)+strings.Repeat("░", 10)),
		MatchRegexp(`slices\s+\| 4/4\s+\| `+strings.Repeat("█", 20)),
		MatchRegexp(`maps\s+\| 0/3\s+\| `+strings.Repeat("░", 20)),
		MatchRegexp(`\(other\)\s+\| 1/2\s+\| `),
	))
}

func TestPrintTopicProgressWithNoExercises(t *testing.T) {
	g := NewWithT(t)
	var out bytes.Buffer

	ui.PrintTopicProgress(&out, []exercises.TopicProgress{{Topic: "empty"}})

	g.Expect(out.String()).To(MatchRegexp(`empty\s+\| 0/0\s+\| ` + strings.Repeat("░", 20)))
}
