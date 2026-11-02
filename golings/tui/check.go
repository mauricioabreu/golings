package tui

import (
	"context"
	"errors"
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/cleonte/golings/golings/exercises"
)

// checkWorkers is how many exercises are built and run at the same time while
// checking all of them.
const checkWorkers = 4

type checkStatus int

const (
	checkWaiting checkStatus = iota
	checkRunning
	checkPassed
	checkFailed
	checkSkipped
)

type checkItem struct {
	exercise exercises.Exercise
	status   checkStatus
	result   exercises.Result
	err      error
}

// next moves on from an exercise that works: it takes out the "I AM NOT DONE"
// marker, which is what the learner would do by hand, and runs the next one.
func (m *Model) next() tea.Cmd {
	switch {
	case !m.hasExercise():
		return nil
	case m.status != statusFinished || m.runErr != nil:
		m.notice = "Get " + m.exercise.Name + " to work first, then n moves on to the next exercise."
		return nil
	}

	if err := m.exercise.MarkDone(); err != nil && !errors.Is(err, exercises.ErrNotMarked) {
		m.notice = "Could not mark " + m.exercise.Name + " as done: " + err.Error()
		return nil
	}

	name := m.exercise.Name
	m.pinned = ""
	cmd := m.startRun()
	m.notice = name + " is done."
	return cmd
}

// startCheck runs every exercise, a few at a time, without changing any file.
func (m *Model) startCheck() tea.Cmd {
	if m.checkCancel != nil {
		m.notice = "Already checking the exercises."
		return nil
	}

	all, err := exercises.List(m.cfg.InfoFile)
	if err != nil || len(all) == 0 {
		m.notice = "There are no exercises to check."
		return nil
	}

	m.stopRun()
	if m.status == statusRunning {
		m.status = statusStarting
	}
	m.checkID++
	ctx, cancel := context.WithCancel(context.Background())
	m.checkCtx, m.checkCancel = ctx, cancel
	m.checkItems = make([]checkItem, len(all))
	for i, exercise := range all {
		m.checkItems[i] = checkItem{exercise: exercise}
	}
	m.checkNext, m.checkActive = 0, 0
	m.mode, m.scroll = modeCheck, 0

	var cmds []tea.Cmd
	for range checkWorkers {
		if cmd := m.startNextCheck(); cmd != nil {
			cmds = append(cmds, cmd)
		}
	}
	return tea.Batch(cmds...)
}

// startNextCheck starts the next exercise that has not been started yet.
func (m *Model) startNextCheck() tea.Cmd {
	if m.checkNext >= len(m.checkItems) {
		return nil
	}
	index := m.checkNext
	m.checkNext++
	m.checkItems[index].status = checkRunning
	m.checkActive++

	exercise, id, run, ctx := m.checkItems[index].exercise, m.checkID, m.cfg.Run, m.checkCtx
	return func() tea.Msg {
		result, err := run(ctx, exercise)
		return checkDoneMsg{id: id, index: index, result: result, err: err}
	}
}

func (m *Model) finishCheck(msg checkDoneMsg) tea.Cmd {
	if msg.id != m.checkID || m.checkCancel == nil || msg.index >= len(m.checkItems) {
		return nil
	}

	item := &m.checkItems[msg.index]
	item.result, item.err = msg.result, msg.err
	switch {
	case errors.Is(msg.err, exercises.ErrGoVersionTooOld):
		item.status = checkSkipped
	case msg.err != nil:
		item.status = checkFailed
	default:
		item.status = checkPassed
	}
	m.checkActive--

	if cmd := m.startNextCheck(); cmd != nil {
		return cmd
	}
	if m.checkActive == 0 {
		m.checkCancel()
		m.checkCancel = nil
	}
	return nil
}

// stopCheck cancels a check that is still going.
func (m *Model) stopCheck() {
	if m.checkCancel != nil {
		m.checkCancel()
		m.checkCancel = nil
		m.checkID++
	}
}

// checking reports whether exercises are still being checked.
func (m *Model) checking() bool {
	return m.checkCancel != nil
}

func (m *Model) checkLines() []string {
	var passed, failed, skipped, done int
	for _, item := range m.checkItems {
		switch item.status {
		case checkPassed:
			passed++
		case checkFailed:
			failed++
		case checkSkipped:
			skipped++
		}
		if item.status >= checkPassed {
			done++
		}
	}

	title := fmt.Sprintf("Checked %d exercises: %d work, %d do not, %d skipped", len(m.checkItems), passed, failed, skipped)
	if m.checking() {
		title = fmt.Sprintf("Checking the exercises… %d/%d", done, len(m.checkItems))
	}
	lines := []string{titleStyle.Render(title), ""}

	for _, item := range m.checkItems {
		lines = append(lines, m.checkLine(item))
	}
	return lines
}

func (m *Model) checkLine(item checkItem) string {
	name := item.exercise.Name
	switch item.status {
	case checkWaiting:
		return dimStyle.Render("· " + name)
	case checkRunning:
		return dimStyle.Render("… " + name)
	case checkSkipped:
		return warnStyle.Render(fmt.Sprintf("• %s  needs Go %s", name, item.exercise.GoVersion))
	case checkPassed:
		if item.exercise.State() == exercises.Pending {
			return okStyle.Render("✔ "+name) + dimStyle.Render("  works, the I AM NOT DONE marker is still there")
		}
		return okStyle.Render("✔ " + name)
	}

	// A failure: say how it failed, and shout if it was finished before.
	how := item.result.Failure().Kind.String()
	if item.exercise.State() == exercises.Done {
		return warnStyle.Render(fmt.Sprintf("✗ %s  was finished, but now %s", name, how))
	}
	return errStyle.Render("✗ "+name) + dimStyle.Render("  "+how)
}
