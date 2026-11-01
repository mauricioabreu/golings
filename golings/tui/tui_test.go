package tui

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
	. "github.com/onsi/gomega"
)

// workspace writes exercises/<topic>/<name>/main.go files ("topic/name",
// a leading "!" leaves it unsolved) and returns the path of its info file.
func workspace(t *testing.T, specs ...string) string {
	t.Helper()
	root := t.TempDir()
	var info strings.Builder
	for _, spec := range specs {
		unsolved := strings.HasPrefix(spec, "!")
		spec = strings.TrimPrefix(spec, "!")
		file := filepath.Join(root, "exercises", spec, "main.go")
		if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
			t.Fatal(err)
		}
		source := "package main\n\nfunc main() {}\n"
		if unsolved {
			source = "// I AM NOT DONE\n" + source
		}
		if err := os.WriteFile(file, []byte(source), 0o644); err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(&info, "[[exercises]]\nname = %q\npath = %q\nmode = \"compile\"\nhint = \"\"\n\n", filepath.Base(spec), file)
	}
	infoFile := filepath.Join(root, "info.toml")
	if err := os.WriteFile(infoFile, []byte(info.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	return infoFile
}

func noColor(t *testing.T) {
	t.Helper()
	prev := color.NoColor
	color.NoColor = true
	t.Cleanup(func() { color.NoColor = prev })
}

// fakeRuns stands in for the real runner and records what it was asked to do.
type fakeRuns struct {
	mu       sync.Mutex
	started  []string
	contexts []context.Context
	result   func(ex exercises.Exercise) (exercises.Result, error)
	// block makes a run wait until it is cancelled, like an exercise that is
	// still going.
	block bool
}

func (f *fakeRuns) run(ctx context.Context, ex exercises.Exercise) (exercises.Result, error) {
	f.mu.Lock()
	f.started = append(f.started, ex.Name)
	f.contexts = append(f.contexts, ctx)
	result, block := f.result, f.block
	f.mu.Unlock()

	if block {
		<-ctx.Done()
		return exercises.Result{Exercise: ex}, ctx.Err()
	}

	if result == nil {
		return exercises.Result{Exercise: ex}, nil
	}
	return result(ex)
}

// exec runs a command the way the Bubble Tea runtime would and returns its message.
func exec(t *testing.T, cmd tea.Cmd) tea.Msg {
	t.Helper()
	if cmd == nil {
		return nil
	}
	return cmd()
}

// start sends the start message and returns the model, the pending run command
// (run synchronously by the caller) and the model after the run finished.
func start(t *testing.T, cfg Config) (*Model, tea.Cmd) {
	t.Helper()
	if cfg.Changes == nil {
		// With a channel to wait on, the start command is always a batch of two.
		cfg.Changes = make(chan []string)
	}
	m := New(cfg)
	_, cmd := m.Update(startMsg{})
	return m, cmd
}

// runCmds executes a command returned by Update, and the commands of a batch,
// feeding finished runs back to the model. A command that is still blocked
// after a moment is the one waiting for file changes and is left alone.
func runCmds(t *testing.T, m *Model, cmd tea.Cmd) {
	t.Helper()
	if cmd == nil {
		return
	}
	msg := settle(cmd)
	if batch, ok := msg.(tea.BatchMsg); ok {
		for _, c := range batch {
			if c != nil {
				runCmds(t, m, c)
			}
		}
		return
	}
	if _, isRun := msg.(runDoneMsg); isRun {
		m.Update(msg)
	}
}

// settle runs cmd and returns its message, or nil if it is still waiting.
func settle(cmd tea.Cmd) tea.Msg {
	done := make(chan tea.Msg, 1)
	go func() { done <- cmd() }()
	select {
	case msg := <-done:
		return msg
	case <-time.After(300 * time.Millisecond):
		return nil
	}
}

func TestShowsProgressTopicAndTheRunningExercise(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := workspace(t, "a/a1", "!a/a2", "!b/b1")
	fake := &fakeRuns{}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})
	m.Update(tea.WindowSizeMsg{Width: 80, Height: 20})

	view := m.View().Content

	g.Expect(cmd).NotTo(BeNil())
	g.Expect(view).To(ContainSubstring("golings"))
	g.Expect(view).To(ContainSubstring("1/3 (33%)"))
	g.Expect(view).To(ContainSubstring("Topic: a (1/2)"))
	g.Expect(view).To(ContainSubstring("a2"))
	g.Expect(view).To(ContainSubstring("Running a2"))
	g.Expect(m.View().AltScreen).To(BeTrue())
}

func TestShowsTheFailureOfTheExercise(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := workspace(t, "!a/a1")
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) {
		return exercises.Result{Exercise: ex, BuildFailed: true, Err: ex.Path + ":4:6: undefined: nope\n"}, errors.New("exit status 1")
	}}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})

	runCmds(t, m, cmd)

	view := m.View().Content
	g.Expect(view).To(ContainSubstring("Failed to compile the exercise"))
	g.Expect(view).To(ContainSubstring("undefined: nope"))
	g.Expect(view).To(ContainSubstring("func main() {}"), "the offending source line is shown")
}

func TestAWorkingExerciseThatIsStillPendingAsksToRemoveTheMarker(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := workspace(t, "!a/a1")
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) {
		return exercises.Result{Exercise: ex, Out: "hello from the exercise\n"}, nil
	}}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})

	runCmds(t, m, cmd)

	view := m.View().Content
	g.Expect(view).To(ContainSubstring("a1 works!"))
	g.Expect(view).To(ContainSubstring("hello from the exercise"))
	g.Expect(view).To(ContainSubstring("I AM NOT DONE"))
}

func TestASaveCancelsTheRunInProgress(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := workspace(t, "!a/a1")
	fake := &fakeRuns{block: true}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})

	// The first run is going, blocked on its context.
	finished := make(chan tea.Msg, 1)
	batch := cmd().(tea.BatchMsg)
	go func() { finished <- batch[0]() }()
	g.Eventually(func() int { fake.mu.Lock(); defer fake.mu.Unlock(); return len(fake.started) }).Should(Equal(1))
	firstID := m.runID

	m.Update(changesMsg{files: []string{"x.go"}})

	g.Expect(m.runID).To(Equal(firstID+1), "a new run replaces the old one")
	var late tea.Msg
	g.Eventually(finished, 2*time.Second).Should(Receive(&late), "the first run was cancelled")
	g.Expect(late).To(BeAssignableToTypeOf(runDoneMsg{}))

	// Its late result must not overwrite the screen.
	m.Update(late)
	g.Expect(m.status).To(Equal(statusRunning))
}

func TestQuitStopsTheRunAndExits(t *testing.T) {
	for _, key := range []tea.KeyPressMsg{{Code: 'q', Text: "q"}, {Code: 'c', Mod: tea.ModCtrl}} {
		t.Run(key.String(), func(t *testing.T) {
			g := NewWithT(t)
			m, _ := start(t, Config{InfoFile: workspace(t, "!a/a1"), Run: (&fakeRuns{block: true}).run})
			before := m.runID

			_, cmd := m.Update(key)

			g.Expect(exec(t, cmd)).To(BeAssignableToTypeOf(tea.QuitMsg{}))
			g.Expect(m.runID).To(Equal(before+1), "results of the stopped run are ignored")
		})
	}
}

func TestOtherKeysDoNothing(t *testing.T) {
	g := NewWithT(t)
	m, _ := start(t, Config{InfoFile: workspace(t, "!a/a1"), Run: (&fakeRuns{}).run})

	_, cmd := m.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})

	g.Expect(cmd).To(BeNil())
}

func TestWhenEverythingIsDone(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	fake := &fakeRuns{}
	m, cmd := start(t, Config{InfoFile: workspace(t, "a/a1", "a/a2"), Run: fake.run})

	runCmds(t, m, cmd)

	g.Expect(m.View().Content).To(ContainSubstring("completed all the exercises"))
	g.Expect(m.View().Content).To(ContainSubstring("2/2 (100%)"))
	g.Expect(fake.started).To(BeEmpty(), "there is nothing left to run")
}

func TestWhenThereIsNothingToRun(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := filepath.Join(t.TempDir(), "info.toml")
	g.Expect(os.WriteFile(info, nil, 0o644)).To(Succeed())
	m, _ := start(t, Config{InfoFile: info, Run: (&fakeRuns{}).run})

	g.Expect(m.View().Content).To(ContainSubstring("no exercises to run"))
	g.Expect(m.View().Content).NotTo(ContainSubstring("NaN"))
}

func TestWhenTheInfoFileCannotBeRead(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	m, _ := start(t, Config{InfoFile: filepath.Join(t.TempDir(), "missing.toml"), Run: (&fakeRuns{}).run})

	view := m.View().Content

	g.Expect(view).To(ContainSubstring("Failed to read the exercises"))
	g.Expect(view).To(ContainSubstring("golings init"))
	g.Expect(view).NotTo(ContainSubstring("Congratulations"))
}

func TestLongOutputIsCutToTheHeight(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	long := strings.Repeat("a line of output\n", 100)
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) {
		return exercises.Result{Exercise: ex, Out: long}, nil
	}}
	m, cmd := start(t, Config{InfoFile: workspace(t, "!a/a1"), Run: fake.run})
	m.Update(tea.WindowSizeMsg{Width: 60, Height: 12})
	runCmds(t, m, cmd)

	view := m.View().Content

	g.Expect(strings.Count(view, "\n") + 1).To(BeNumerically("<=", 12))
	g.Expect(view).To(ContainSubstring("more lines"))
	g.Expect(view).To(ContainSubstring("q quit"), "the footer stays visible")
}

func TestKeepsWorkingWhenTheWatcherStops(t *testing.T) {
	g := NewWithT(t)
	m, _ := start(t, Config{InfoFile: workspace(t, "!a/a1"), Run: (&fakeRuns{}).run})

	_, cmd := m.Update(changesClosedMsg{})

	g.Expect(cmd).To(BeNil())
}

func TestWaitsForTheNextSave(t *testing.T) {
	g := NewWithT(t)
	changes := make(chan []string, 1)
	m := New(Config{InfoFile: workspace(t, "!a/a1"), Changes: changes})

	wait := m.waitForChanges()
	changes <- []string{"a.go"}
	g.Expect(exec(t, wait)).To(Equal(changesMsg{files: []string{"a.go"}}))

	close(changes)
	g.Expect(exec(t, m.waitForChanges())).To(Equal(changesClosedMsg{}))
}

// syncBuffer collects what the program draws while the test reads it.
type syncBuffer struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

// The whole program, with the real event loop but a fake terminal: it draws the
// first run, redraws after a save, and quits on q.
func TestProgramRunsAndQuits(t *testing.T) {
	g := NewWithT(t)
	noColor(t)

	var runs int
	var mu sync.Mutex
	run := func(ctx context.Context, ex exercises.Exercise) (exercises.Result, error) {
		mu.Lock()
		runs++
		n := runs
		mu.Unlock()
		return exercises.Result{Exercise: ex, Out: fmt.Sprintf("output of run %d\n", n)}, nil
	}
	changes := make(chan []string)
	input, typed := newPipe()
	out := &syncBuffer{}
	program := tea.NewProgram(
		New(Config{InfoFile: workspace(t, "!a/a1"), Changes: changes, Run: run}),
		tea.WithInput(input), tea.WithOutput(out), tea.WithWindowSize(80, 24), tea.WithoutSignalHandler(),
	)

	type outcome struct {
		model tea.Model
		err   error
	}
	finished := make(chan outcome, 1)
	go func() { model, err := program.Run(); finished <- outcome{model, err} }()

	g.Eventually(out.String, 5*time.Second).Should(ContainSubstring("output of run 1"))
	changes <- []string{"exercises/a/a1/main.go"}
	// The renderer only redraws what changed, so look at the model, not the bytes.
	g.Eventually(func() int { mu.Lock(); defer mu.Unlock(); return runs }, 5*time.Second).Should(Equal(2))

	_, err := typed.Write([]byte("q"))
	g.Expect(err).NotTo(HaveOccurred())
	var result outcome
	g.Eventually(finished, 5*time.Second).Should(Receive(&result))
	g.Expect(result.err).NotTo(HaveOccurred())
	g.Expect(result.model.(*Model).View().Content).To(ContainSubstring("output of run 2"))
}

func newPipe() (*io.PipeReader, *io.PipeWriter) { return io.Pipe() }
