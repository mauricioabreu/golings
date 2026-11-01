package ui

import (
	"io"
	"strconv"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/jedib0t/go-pretty/v6/table"
)

const barWidth = 20

// PrintTopicProgress prints one row per topic with a progress bar.
func PrintTopicProgress(o io.Writer, topics []exercises.TopicProgress) {
	t := table.NewWriter()
	t.SetOutputMirror(o)
	t.AppendHeader(table.Row{"Topic", "Done", "Progress"})
	for _, topic := range topics {
		name := topic.Topic
		if name == "" {
			name = "(other)"
		}
		t.AppendRow(table.Row{name, formatDone(topic), bar(topic.Done, topic.Total)})
	}
	t.Render()
}

func formatDone(topic exercises.TopicProgress) string {
	return strconv.Itoa(topic.Done) + "/" + strconv.Itoa(topic.Total)
}

func bar(done, total int) string {
	if total == 0 {
		return strings.Repeat("░", barWidth)
	}
	filled := done * barWidth / total
	return strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
}
