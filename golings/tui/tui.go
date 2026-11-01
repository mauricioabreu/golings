// Package tui is the full-screen terminal interface of `golings watch`: the
// progress at the top, the result of the current exercise below, refreshed
// every time a file is saved.
package tui

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
)

// RunFunc builds and runs an exercise. It must give up when ctx is cancelled.
type RunFunc func(ctx context.Context, ex exercises.Exercise) (exercises.Result, error)

type Config struct {
	InfoFile string
	// Changes delivers the files that were saved, see the watcher package.
	Changes <-chan []string
	// Run defaults to running the exercise with exercises.DefaultTimeout.
	Run RunFunc
}

type status int

const (
	statusStarting status = iota
	statusRunning
	statusFinished
	statusAllDone
	statusNothingToRun
	statusError
)

type Model struct {
	cfg    Config
	width  int
	height int

	status   status
	message  string // for statusError
	done     int
	total    int
	skipped  int
	topic    string // e.g. "Topic: slices (2/5)", empty when the exercise has none
	exercise exercises.Exercise

	// The run in progress. Results of earlier runs carry an older id and are
	// ignored: a save during a run cancels it and starts a new one.
	runID  int
	cancel context.CancelFunc
	result exercises.Result
	runErr error
}

type startMsg struct{}

type changesMsg struct{ files []string }

type changesClosedMsg struct{}

type runDoneMsg struct {
	id     int
	result exercises.Result
	err    error
}

func New(cfg Config) *Model {
	if cfg.Run == nil {
		cfg.Run = func(ctx context.Context, ex exercises.Exercise) (exercises.Result, error) {
			return ex.RunContext(ctx, exercises.DefaultTimeout)
		}
	}
	return &Model{cfg: cfg, width: 80, height: 24}
}

// Run shows the interface until the learner quits or ctx ends.
func Run(ctx context.Context, cfg Config) error {
	_, err := tea.NewProgram(New(cfg), tea.WithContext(ctx)).Run()
	if errors.Is(err, tea.ErrProgramKilled) && ctx.Err() != nil {
		return nil
	}
	return err
}

func (m *Model) Init() tea.Cmd {
	return func() tea.Msg { return startMsg{} }
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.stopRun()
			return m, tea.Quit
		}

	case startMsg:
		return m, tea.Batch(m.startRun(), m.waitForChanges())

	case changesMsg:
		return m, tea.Batch(m.startRun(), m.waitForChanges())

	case changesClosedMsg:
		// The watcher stopped; whatever is on screen stays, and q still works.

	case runDoneMsg:
		if msg.id == m.runID {
			m.status, m.result, m.runErr = statusFinished, msg.result, msg.err
		}
	}
	return m, nil
}

func (m *Model) stopRun() {
	if m.cancel != nil {
		m.cancel()
		m.cancel = nil
	}
	m.runID++
}

// startRun cancels the run in progress, works out which exercise is next and
// starts running it.
func (m *Model) startRun() tea.Cmd {
	m.stopRun()

	_, done, total, err := exercises.Progress(m.cfg.InfoFile)
	if err != nil {
		m.status, m.message = statusError, fmt.Sprintf("Failed to read the exercises: %s", err)
		return nil
	}
	m.done, m.total = done, total
	if skipped, err := exercises.Skipped(m.cfg.InfoFile); err == nil {
		m.skipped = len(skipped)
	}

	switch {
	case total == 0:
		m.status = statusNothingToRun
		return nil
	case done == total:
		m.status = statusAllDone
		return nil
	}

	exercise, err := exercises.NextPending(m.cfg.InfoFile)
	if err != nil {
		m.status, m.message = statusError, "Failed to find the next exercise"
		return nil
	}
	m.exercise = exercise
	m.topic = m.topicLine(exercise)

	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	m.status, m.result, m.runErr = statusRunning, exercises.Result{}, nil

	id, run := m.runID, m.cfg.Run
	return func() tea.Msg {
		result, err := run(ctx, exercise)
		return runDoneMsg{id: id, result: result, err: err}
	}
}

func (m *Model) topicLine(exercise exercises.Exercise) string {
	if exercise.Topic() == "" {
		return ""
	}
	topics, err := exercises.ProgressByTopic(m.cfg.InfoFile)
	if err != nil {
		return ""
	}
	for _, t := range topics {
		if t.Topic == exercise.Topic() {
			return fmt.Sprintf("Topic: %s (%d/%d)", t.Topic, t.Done, t.Total)
		}
	}
	return ""
}

func (m *Model) waitForChanges() tea.Cmd {
	changes := m.cfg.Changes
	if changes == nil {
		return nil
	}
	return func() tea.Msg {
		files, ok := <-changes
		if !ok {
			return changesClosedMsg{}
		}
		return changesMsg{files: files}
	}
}

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("6"))
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	okStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))
	warnStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	errStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	headerLines = 3 // title, exercise, rule
	footerLines = 1
)

func (m *Model) View() tea.View {
	var b strings.Builder
	b.WriteString(m.header())
	b.WriteString(m.body())

	view := tea.NewView(b.String())
	view.AltScreen = true
	view.WindowTitle = "golings"
	return view
}

func (m *Model) header() string {
	title := titleStyle.Render("golings")
	if m.total > 0 {
		title += fmt.Sprintf("  %d/%d (%.0f%%)", m.done, m.total, 100*float64(m.done)/float64(m.total))
	}
	if m.topic != "" {
		title += "  " + dimStyle.Render(m.topic)
	}
	if m.skipped > 0 {
		title += "  " + warnStyle.Render(fmt.Sprintf("%d skipped (newer Go needed)", m.skipped))
	}

	current := ""
	switch m.status {
	case statusRunning, statusFinished:
		current = fmt.Sprintf("%s  %s", m.exercise.Name, dimStyle.Render(m.exercise.Path))
	}

	rule := dimStyle.Render(strings.Repeat("─", max(m.width, 1)))
	return title + "\n" + current + "\n" + rule + "\n"
}

func (m *Model) body() string {
	var lines []string

	switch m.status {
	case statusStarting:
		lines = []string{"Starting…"}
	case statusRunning:
		lines = []string{dimStyle.Render("Running " + m.exercise.Name + "…")}
	case statusAllDone:
		lines = []string{
			okStyle.Render("Congratulations! You have completed all the exercises."),
			"",
			"If you enjoyed it, please give the repository a star:",
			"https://github.com/cleonte/golings",
		}
	case statusNothingToRun:
		lines = []string{warnStyle.Render("There are no exercises to run."), "Try `golings doctor` to check your setup."}
	case statusError:
		lines = []string{errStyle.Render(m.message), "Run golings from the folder that holds info.toml, create one with `golings init`, or try `golings doctor`."}
	case statusFinished:
		lines = m.resultLines()
	}

	available := max(m.height-headerLines-footerLines, 1)
	if len(lines) > available {
		hidden := len(lines) - (available - 1)
		lines = append(lines[:available-1], dimStyle.Render(fmt.Sprintf("… %d more lines", hidden)))
	}
	for len(lines) < available {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n") + "\n" + dimStyle.Render("q quit")
}

func (m *Model) resultLines() []string {
	var lines []string

	switch {
	case errors.Is(m.runErr, exercises.ErrGoVersionTooOld):
		lines = append(lines, warnStyle.Render(m.runErr.Error()))

	case m.runErr != nil:
		var out bytes.Buffer
		ui.PrintFailure(&out, m.exercise, m.result.Failure())
		lines = append(lines, strings.Split(strings.TrimRight(out.String(), "\n"), "\n")...)
		lines = append(lines, "", dimStyle.Render("Stuck? Type `golings hint "+m.exercise.Name+"` in another terminal."))

	default:
		lines = append(lines, okStyle.Render("✔ "+m.exercise.Name+" works!"), "")
		if output := strings.TrimRight(m.result.Out, "\n"); output != "" {
			lines = append(lines, "Here is the output of your program:", "")
			lines = append(lines, strings.Split(output, "\n")...)
			lines = append(lines, "")
		}
		if m.exercise.State() == exercises.Pending {
			lines = append(lines, warnStyle.Render("Remove the `// I AM NOT DONE` marker from the file to move on to the next exercise."))
		}
	}
	return lines
}
