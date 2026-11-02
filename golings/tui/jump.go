package tui

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/cleonte/golings/golings/exercises"
)

// jumpItem is one row of the list of exercises.
type jumpItem struct {
	exercise exercises.Exercise
	done     bool
	skipped  bool // the local Go is too old to run it
}

// FilterValue lets the learner find an exercise by its name or its topic.
func (i jumpItem) FilterValue() string {
	return i.exercise.Name + " " + i.exercise.Topic()
}

// jumpDelegate draws one line per exercise.
type jumpDelegate struct{}

func (jumpDelegate) Height() int                         { return 1 }
func (jumpDelegate) Spacing() int                        { return 0 }
func (jumpDelegate) Update(tea.Msg, *list.Model) tea.Cmd { return nil }

func (jumpDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	it, ok := item.(jumpItem)
	if !ok {
		return
	}

	mark := dimStyle.Render("·")
	switch {
	case it.skipped:
		mark = warnStyle.Render("•")
	case it.done:
		mark = okStyle.Render("✔")
	}

	cursor := "  "
	name := it.exercise.Name
	if index == m.Index() {
		cursor = titleStyle.Render("> ")
		name = titleStyle.Render(name)
	}

	line := fmt.Sprintf("%s%s %s", cursor, mark, name)
	if it.skipped {
		line += warnStyle.Render("  needs Go " + it.exercise.GoVersion)
	} else if topic := it.exercise.Topic(); topic != "" {
		line += dimStyle.Render("  " + topic)
	}
	_, _ = fmt.Fprint(w, lipgloss.NewStyle().MaxWidth(max(m.Width(), 1)).Render(line))
}

// openList shows every exercise, with the current one selected.
func (m *Model) openList() {
	all, err := exercises.List(m.cfg.InfoFile)
	if err != nil || len(all) == 0 {
		m.notice = "There are no exercises to list."
		return
	}

	skipped := map[string]bool{}
	if s, err := exercises.Skipped(m.cfg.InfoFile); err == nil {
		for _, e := range s {
			skipped[e.Name] = true
		}
	}

	items := make([]list.Item, len(all))
	current := 0
	for i, e := range all {
		items[i] = jumpItem{exercise: e, done: e.State() == exercises.Done, skipped: skipped[e.Name]}
		if m.hasExercise() && e.Name == m.exercise.Name {
			current = i
		}
	}

	l := list.New(items, jumpDelegate{}, m.width, m.available())
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(true)
	l.DisableQuitKeybindings()
	l.Select(current)

	m.list = l
	m.mode, m.scroll = modeList, 0
}

// listKey handles a key while the list is open. The list gets most keys, but
// not q (unless a filter is being typed, where it is just a letter) and not
// enter, which opens the exercise under the cursor.
func (m *Model) listKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	key := msg.String()
	m.notice = ""

	switch {
	case key == "ctrl+c", key == "q" && !m.list.SettingFilter():
		m.stopRun()
		m.stopCheck()
		return m, tea.Quit

	case key == "esc" && !m.list.SettingFilter() && !m.list.IsFiltered():
		m.mode = modeResult
		return m, nil

	case key == "enter" && !m.list.SettingFilter():
		if it, ok := m.list.SelectedItem().(jumpItem); ok {
			m.pinned = it.exercise.Name
			return m, m.startRun()
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) listLines() []string {
	return strings.Split(m.list.View(), "\n")
}
