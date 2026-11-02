// Package tui is the full-screen terminal interface of `golings watch`: the
// progress at the top, the result of the current exercise below, refreshed
// every time a file is saved.
package tui

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	"github.com/cleonte/golings/golings/workspace"
)

// RunFunc builds and runs an exercise. It must give up when ctx is cancelled.
type RunFunc func(ctx context.Context, ex exercises.Exercise) (exercises.Result, error)

type Config struct {
	InfoFile string
	// Changes delivers the files that were saved, see the watcher package.
	Changes <-chan []string
	// Run defaults to running the exercise with exercises.DefaultTimeout.
	Run RunFunc
	// Source holds the original exercises, used to reset one. Root is the folder
	// the exercise paths of the info file are relative to. Without a Source the
	// reset key says it is not available.
	Source fs.FS
	Root   string
	// Solutions holds the solutions shown on request, laid out like
	// exercises/. Without it the solution key says it is not available.
	Solutions fs.FS
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

// mode is what the lower part of the screen is showing.
type mode int

const (
	modeResult mode = iota // the result of the run
	modeHint
	modeSolution
	modeConfirmSolution
	modeConfirmReset
	modeCheck // the result of checking every exercise
)

type Model struct {
	cfg    Config
	width  int
	height int

	mode   mode
	scroll int    // lines scrolled off the top of the body
	notice string // one line shown instead of the key help, e.g. "Reset a1."

	status   status
	message  string // for statusError
	done     int
	total    int
	skipped  int
	topic    string // e.g. "Topic: slices (2/5)", empty when the exercise has none
	exercise exercises.Exercise

	// Checking every exercise (the c key). Like runs, the results of a check
	// that was cancelled carry an older id and are ignored.
	checkID     int
	checkCtx    context.Context
	checkCancel context.CancelFunc
	checkItems  []checkItem
	checkNext   int // index of the next exercise to start
	checkActive int // how many are running now

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

type checkDoneMsg struct {
	id     int
	index  int
	result exercises.Result
	err    error
}

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
		return m.handleKey(msg)

	case startMsg:
		return m, tea.Batch(m.startRun(), m.waitForChanges())

	case changesMsg:
		if m.checking() {
			// A save should not throw away a check that is half way (golings
			// itself saves a file when n takes out a marker). Leaving the check
			// with esc runs the exercise again.
			return m, m.waitForChanges()
		}
		return m, tea.Batch(m.startRun(), m.waitForChanges())

	case changesClosedMsg:
		// The watcher stopped; whatever is on screen stays, and q still works.

	case checkDoneMsg:
		return m, m.finishCheck(msg)

	case runDoneMsg:
		if msg.id == m.runID {
			m.status, m.result, m.runErr = statusFinished, msg.result, msg.err
		}
	}
	return m, nil
}

func (m *Model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	key := msg.String()
	if key == "q" || key == "ctrl+c" {
		m.stopRun()
		m.stopCheck()
		return m, tea.Quit
	}
	m.notice = ""

	// A question is open: only yes or no answer it.
	if m.mode == modeConfirmSolution || m.mode == modeConfirmReset {
		pending := m.mode
		m.mode, m.scroll = modeResult, 0
		if key != "y" {
			return m, nil
		}
		if pending == modeConfirmSolution {
			m.showSolution()
			return m, nil
		}
		return m, m.reset()
	}

	// While checking, only leaving, scrolling and a second c (which says it is
	// already going) make sense.
	if m.mode == modeCheck && !slices.Contains([]string{"esc", "c", "down", "j", "up", "k", "pgdown", "space", "pgup"}, key) {
		return m, nil
	}

	switch key {
	case "esc":
		if m.mode == modeCheck {
			// The run in progress was stopped for the check: run it again.
			return m, m.startRun()
		}
		m.mode, m.scroll = modeResult, 0
	case "h":
		m.toggle(modeHint)
	case "s":
		m.askToShowSolution()
	case "r":
		m.askToReset()
	case "n":
		return m, m.next()
	case "c":
		return m, m.startCheck()
	case "down", "j":
		m.scroll = min(m.scroll+1, m.maxScroll())
	case "up", "k":
		m.scroll = max(m.scroll-1, 0)
	case "pgdown", "space":
		m.scroll = min(m.scroll+m.available(), m.maxScroll())
	case "pgup":
		m.scroll = max(m.scroll-m.available(), 0)
	}
	return m, nil
}

// hasExercise reports whether there is a current exercise to act on.
func (m *Model) hasExercise() bool {
	return m.status == statusRunning || m.status == statusFinished
}

func (m *Model) toggle(target mode) {
	if !m.hasExercise() {
		return
	}
	m.scroll = 0
	if m.mode == target {
		m.mode = modeResult
		return
	}
	m.mode = target
}

func (m *Model) askToShowSolution() {
	switch {
	case !m.hasExercise():
	case m.cfg.Solutions == nil:
		m.notice = "Solutions are not available here."
	default:
		m.mode, m.scroll = modeConfirmSolution, 0
	}
}

func (m *Model) showSolution() {
	if _, err := fs.Stat(m.cfg.Solutions, m.exercise.SolutionPath()); err != nil {
		m.notice = "No solution available for " + m.exercise.Name + " yet."
		return
	}
	m.mode, m.scroll = modeSolution, 0
}

func (m *Model) askToReset() {
	switch {
	case !m.hasExercise():
	case m.cfg.Source == nil:
		m.notice = "Reset is not available here."
	default:
		m.mode, m.scroll = modeConfirmReset, 0
	}
}

// reset puts the original exercise back and runs it again.
func (m *Model) reset() tea.Cmd {
	root := m.cfg.Root
	if root == "" {
		root = "."
	}
	path, err := workspace.Reset(m.cfg.Source, root, m.exercise.Name)
	if err != nil {
		m.notice = "Reset failed: " + err.Error()
		return nil
	}
	cmd := m.startRun()
	m.notice = "Reset " + path + " to its original state."
	return cmd
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
	m.stopCheck()
	m.mode, m.scroll = modeResult, 0

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

// available is the number of lines the body can use.
func (m *Model) available() int {
	return max(m.height-headerLines-footerLines, 1)
}

// maxScroll is how far the body can be scrolled before running out of lines.
func (m *Model) maxScroll() int {
	return max(len(m.lines())-m.available(), 0)
}

// lines is everything the body would show, before cutting it to the screen.
func (m *Model) lines() []string {
	switch m.mode {
	case modeCheck:
		return m.checkLines()
	case modeHint:
		return m.hintLines()
	case modeSolution:
		return m.solutionLines()
	case modeConfirmSolution:
		return []string{
			warnStyle.Render("Show the solution of " + m.exercise.Name + "?"),
			"It spoils the exercise: try a bit longer, or read the hint (esc, then h) first.",
			"",
			"y to show it, any other key to go back",
		}
	case modeConfirmReset:
		return []string{
			warnStyle.Render("Reset " + m.exercise.Name + " to its original state?"),
			"This discards your changes to " + m.exercise.Path + ".",
			"",
			"y to reset it, any other key to keep your work",
		}
	}

	switch m.status {
	case statusStarting:
		return []string{"Starting…"}
	case statusRunning:
		return []string{dimStyle.Render("Running " + m.exercise.Name + "…")}
	case statusAllDone:
		return []string{
			okStyle.Render("Congratulations! You have completed all the exercises."),
			"",
			"If you enjoyed it, please give the repository a star:",
			"https://github.com/cleonte/golings",
		}
	case statusNothingToRun:
		return []string{warnStyle.Render("There are no exercises to run."), "Try `golings doctor` to check your setup."}
	case statusError:
		return []string{errStyle.Render(m.message), "Run golings from the folder that holds info.toml, create one with `golings init`, or try `golings doctor`."}
	case statusFinished:
		return m.resultLines()
	}
	return nil
}

func (m *Model) hintLines() []string {
	lines := []string{titleStyle.Render("Hint for " + m.exercise.Name), ""}
	hint := strings.TrimSpace(m.exercise.Hint)
	if hint == "" {
		hint = "There is no hint for this exercise."
	}
	return append(lines, strings.Split(hint, "\n")...)
}

func (m *Model) solutionLines() []string {
	lines := []string{titleStyle.Render("Solution for " + m.exercise.Name), ""}
	source, err := fs.ReadFile(m.cfg.Solutions, m.exercise.SolutionPath())
	if err != nil {
		return append(lines, errStyle.Render(err.Error()))
	}
	source = bytes.ReplaceAll(source, []byte("\t"), []byte("    "))
	return append(lines, strings.Split(strings.TrimRight(string(source), "\n"), "\n")...)
}

func (m *Model) body() string {
	lines := m.lines()
	available := m.available()

	start := min(m.scroll, max(len(lines)-available, 0))
	lines = lines[start:]
	if len(lines) > available {
		hidden := len(lines) - (available - 1)
		lines = append(lines[:available-1], dimStyle.Render(fmt.Sprintf("… %d more lines (↓ to scroll)", hidden)))
	}
	for len(lines) < available {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n") + "\n" + m.footer()
}

func (m *Model) footer() string {
	switch {
	case m.notice != "":
		return warnStyle.Render(m.notice)
	case m.mode == modeConfirmSolution || m.mode == modeConfirmReset:
		return dimStyle.Render("y yes   any other key: no   q quit")
	case m.mode != modeResult:
		return dimStyle.Render("esc back   ↑↓ scroll   q quit")
	case m.hasExercise():
		return dimStyle.Render("h hint   s solution   r reset   n next   c check all   ↑↓ scroll   q quit")
	}
	return dimStyle.Render("q quit")
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
