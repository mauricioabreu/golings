package tui

import (
	"context"
	"errors"
	"os"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/gomega"
)

// pump runs a command, and every command that comes out of feeding its message
// to the model, until nothing is left. A command that is still blocked after a
// moment is the one waiting for file changes and is left alone.
func pump(t *testing.T, m *Model, cmd tea.Cmd) {
	t.Helper()
	if cmd == nil {
		return
	}
	switch msg := settle(cmd).(type) {
	case tea.BatchMsg:
		for _, c := range msg {
			pump(t, m, c)
		}
	case checkDoneMsg, runDoneMsg:
		_, next := m.Update(msg)
		pump(t, m, next)
	}
}

// plain removes the colors from what the view draws.
func plain(view string) string {
	return regexp.MustCompile(`\x1b\[[0-9;]*m`).ReplaceAllString(view, "")
}

func marker(t *testing.T, path string) bool {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return strings.Contains(string(data), "I AM NOT DONE")
}

func passing(ex exercises.Exercise) (exercises.Result, error) {
	return exercises.Result{Exercise: ex, Out: "ok\n"}, nil
}

func failingToCompile(ex exercises.Exercise) (exercises.Result, error) {
	return exercises.Result{Exercise: ex, BuildFailed: true, Err: ex.Path + ":4:1: undefined: x\n"}, errors.New("exit status 1")
}

func TestNMarksTheExerciseDoneAndMovesOn(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := newWorkspace(t, "!a/a1", "!a/a2")
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) { return passing(ex) }}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})
	runCmds(t, m, cmd)
	first := m.exercise
	g.Expect(m.status).To(Equal(statusFinished))
	g.Expect(marker(t, first.Path)).To(BeTrue())

	next := press(m, "n")
	pump(t, m, next)

	g.Expect(marker(t, first.Path)).To(BeFalse(), "the marker is gone from the file")
	g.Expect(m.exercise.Name).To(Equal("a2"), "the next exercise is the one running now")
	g.Expect(fake.started).To(ContainElement("a2"))
	g.Expect(m.notice).To(Equal("a1 is done."))
	g.Expect(m.View().Content).To(ContainSubstring("1/2"), "the progress moved")
}

func TestNDoesNotSkipAnExerciseThatDoesNotWork(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := newWorkspace(t, "!a/a1", "!a/a2")
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) { return failingToCompile(ex) }}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})
	runCmds(t, m, cmd)

	next := press(m, "n")

	g.Expect(next).To(BeNil())
	g.Expect(marker(t, m.exercise.Path)).To(BeTrue(), "the file is untouched")
	g.Expect(m.exercise.Name).To(Equal("a1"))
	g.Expect(m.View().Content).To(ContainSubstring("Get a1 to work first"))
}

func TestNWhileTheExerciseIsStillRunning(t *testing.T) {
	g := NewWithT(t)
	m, _ := start(t, Config{InfoFile: newWorkspace(t, "!a/a1"), Run: (&fakeRuns{block: true}).run})

	next := press(m, "n")

	g.Expect(next).To(BeNil())
	g.Expect(marker(t, m.exercise.Path)).To(BeTrue())
}

func TestNWhenThereIsNothingLeft(t *testing.T) {
	g := NewWithT(t)
	m, _ := start(t, Config{InfoFile: newWorkspace(t, "a/a1"), Run: (&fakeRuns{}).run})
	g.Expect(m.status).To(Equal(statusAllDone))

	g.Expect(press(m, "n")).To(BeNil())
	g.Expect(m.notice).To(BeEmpty())
}

func TestNReportsAFileThatCannotBeChanged(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) { return passing(ex) }}
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, "!a/a1", "!a/a2"), Run: fake.run})
	runCmds(t, m, cmd)
	g.Expect(os.Remove(m.exercise.Path)).To(Succeed())

	next := press(m, "n")

	g.Expect(next).To(BeNil())
	g.Expect(m.notice).To(HavePrefix("Could not mark a1 as done"))
}

// resultsFor gives each exercise the outcome its name says.
func resultsFor(outcomes map[string]func(exercises.Exercise) (exercises.Result, error)) *fakeRuns {
	return &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) {
		if outcome, ok := outcomes[ex.Name]; ok {
			return outcome(ex)
		}
		return passing(ex)
	}}
}

func TestCheckAllShowsWhatEveryExerciseDoes(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	// a1 and a5 are finished (no marker), a2 to a4 are not.
	info := newWorkspace(t, "a/a1", "!a/a2", "!a/a3", "!a/a4", "a/a5")
	fake := resultsFor(map[string]func(exercises.Exercise) (exercises.Result, error){
		"a3": failingToCompile,
		"a4": func(ex exercises.Exercise) (exercises.Result, error) {
			return exercises.Result{Exercise: ex}, exercises.ErrGoVersionTooOld
		},
		"a5": failingToCompile,
	})
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})
	runCmds(t, m, cmd)

	check := press(m, "c")
	pump(t, m, check)

	view := plain(m.View().Content)
	g.Expect(m.checking()).To(BeFalse(), "everything was checked")
	g.Expect(view).To(ContainSubstring("Checked 5 exercises: 2 work, 2 do not, 1 skipped"))
	g.Expect(view).To(MatchRegexp(`✔ a1\n`))
	g.Expect(view).To(ContainSubstring("✔ a2  works, the I AM NOT DONE marker is still there"))
	g.Expect(view).To(ContainSubstring("✗ a3  does not compile"))
	g.Expect(view).To(ContainSubstring("• a4  needs Go"))
	g.Expect(view).To(ContainSubstring("✗ a5  was finished, but now does not compile"))
	g.Expect(m.mode).To(Equal(modeCheck))
}

func TestCheckAllChangesNoFile(t *testing.T) {
	g := NewWithT(t)
	info := newWorkspace(t, "!a/a1", "!a/a2")
	fake := &fakeRuns{}
	m, cmd := start(t, Config{InfoFile: info, Run: fake.run})
	runCmds(t, m, cmd)
	before := marker(t, m.exercise.Path)

	pump(t, m, press(m, "c"))

	g.Expect(marker(t, m.exercise.Path)).To(Equal(before))
	g.Expect(fake.started).To(ContainElements("a1", "a2"))
}

// gatedRuns holds every run until it is released, and counts how many are going.
type gatedRuns struct {
	started atomic.Int32
	mu      sync.Mutex
	gates   map[string]chan struct{}
}

func (g *gatedRuns) gate(name string) chan struct{} {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.gates == nil {
		g.gates = map[string]chan struct{}{}
	}
	if g.gates[name] == nil {
		g.gates[name] = make(chan struct{})
	}
	return g.gates[name]
}

func (g *gatedRuns) run(ctx context.Context, ex exercises.Exercise) (exercises.Result, error) {
	g.started.Add(1)
	select {
	case <-g.gate(ex.Name):
		return exercises.Result{Exercise: ex}, nil
	case <-ctx.Done():
		return exercises.Result{Exercise: ex}, ctx.Err()
	}
}

// bus carries the messages of commands that run in the background to the test.
type bus chan tea.Msg

func (b bus) run(cmd tea.Cmd) {
	if cmd == nil {
		return
	}
	go func() {
		switch msg := cmd().(type) {
		case tea.BatchMsg:
			for _, c := range msg {
				b.run(c)
			}
		default:
			b <- msg
		}
	}()
}

// feed gives the model the next message that arrives, and runs what it returns.
func (b bus) feed(t *testing.T, m *Model) {
	t.Helper()
	select {
	case msg := <-b:
		_, next := m.Update(msg)
		b.run(next)
	case <-time.After(2 * time.Second):
		t.Fatal("no message arrived")
	}
}

func TestCheckAllRunsAFewAtATime(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	gated := &gatedRuns{}
	info := newWorkspace(t, "!a/a1", "!a/a2", "!a/a3", "!a/a4", "!a/a5", "!a/a6")
	m := New(Config{InfoFile: info, Run: gated.run})
	messages := make(bus, 16)

	messages.run(press(m, "c"))

	g.Eventually(gated.started.Load).Should(Equal(int32(checkWorkers)))
	g.Consistently(gated.started.Load, 200*time.Millisecond).Should(Equal(int32(checkWorkers)), "no more than the workers run at once")

	close(gated.gate("a1"))
	messages.feed(t, m)
	g.Eventually(gated.started.Load).Should(Equal(int32(checkWorkers + 1)))
	g.Expect(m.View().Content).To(ContainSubstring("Checking the exercises… 1/6"))

	for _, name := range []string{"a2", "a3", "a4", "a5", "a6"} {
		close(gated.gate(name))
		messages.feed(t, m)
	}

	g.Expect(m.checking()).To(BeFalse())
	g.Expect(m.View().Content).To(ContainSubstring("Checked 6 exercises: 6 work, 0 do not, 0 skipped"))
}

func TestEscapeStopsTheCheckAndRunsTheExerciseAgain(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	gated := &gatedRuns{}
	m := New(Config{InfoFile: newWorkspace(t, "!a/a1", "!a/a2", "!a/a3"), Run: gated.run})
	messages := make(bus, 16)
	messages.run(press(m, "c"))
	g.Eventually(gated.started.Load).Should(Equal(int32(3)))
	id := m.checkID

	cmd := press(m, "esc")

	g.Expect(m.checking()).To(BeFalse())
	g.Expect(m.checkID).To(BeNumerically(">", id))
	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(cmd).NotTo(BeNil(), "the exercise that was shown is run again")
	// the cancelled runs report back late; their results must be ignored
	for range 3 {
		messages.feed(t, m)
	}
	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(m.checkItems[0].status).To(Equal(checkRunning), "a late result of a stopped check changes nothing")
}

func TestASaveDuringTheCheckDoesNotThrowItAway(t *testing.T) {
	g := NewWithT(t)
	gated := &gatedRuns{}
	m := New(Config{InfoFile: newWorkspace(t, "!a/a1", "!a/a2"), Run: gated.run, Changes: make(chan []string)})
	bus(make(chan tea.Msg, 16)).run(press(m, "c"))
	g.Eventually(gated.started.Load).Should(Equal(int32(2)))
	id := m.runID

	_, cmd := m.Update(changesMsg{files: []string{"x.go"}})

	g.Expect(m.checking()).To(BeTrue())
	g.Expect(m.mode).To(Equal(modeCheck))
	g.Expect(m.runID).To(Equal(id), "no new run was started")
	g.Expect(cmd).NotTo(BeNil(), "it keeps waiting for the next save")
}

func TestCheckAllKeys(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	gated := &gatedRuns{}
	m := New(Config{InfoFile: newWorkspace(t, "!a/a1", "!a/a2"), Run: gated.run, Solutions: nil})
	bus(make(chan tea.Msg, 16)).run(press(m, "c"))
	g.Eventually(gated.started.Load).Should(Equal(int32(2)))

	g.Expect(press(m, "c")).To(BeNil())
	g.Expect(m.notice).To(Equal("Already checking the exercises."))

	for _, key := range []string{"h", "s", "r", "n"} {
		g.Expect(press(m, key)).To(BeNil(), key)
		g.Expect(m.mode).To(Equal(modeCheck), "%s does not leave the check", key)
	}

	g.Expect(m.View().Content).To(ContainSubstring("esc back"))
}

func TestQuitCancelsTheCheck(t *testing.T) {
	g := NewWithT(t)
	gated := &gatedRuns{}
	m := New(Config{InfoFile: newWorkspace(t, "!a/a1", "!a/a2"), Run: gated.run})
	bus(make(chan tea.Msg, 16)).run(press(m, "c"))
	g.Eventually(gated.started.Load).Should(Equal(int32(2)))

	quit := press(m, "q")

	g.Expect(exec(t, quit)).To(BeAssignableToTypeOf(tea.QuitMsg{}))
	g.Expect(m.checking()).To(BeFalse())
}

func TestCheckAllWithoutExercises(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	m := New(Config{InfoFile: newWorkspace(t), Run: (&fakeRuns{}).run})

	g.Expect(press(m, "c")).To(BeNil())
	g.Expect(m.notice).To(Equal("There are no exercises to check."))
}

func TestCheckAllScrolls(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	var specs []string
	for i := range 30 {
		specs = append(specs, "!a/a"+strings.Repeat("x", i+1))
	}
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, specs...), Run: (&fakeRuns{}).run})
	runCmds(t, m, cmd)
	m.Update(tea.WindowSizeMsg{Width: 60, Height: 12})
	pump(t, m, press(m, "c"))
	g.Expect(m.View().Content).NotTo(ContainSubstring(strings.Repeat("x", 30)))

	for range 100 {
		press(m, "down")
	}

	g.Expect(m.View().Content).To(ContainSubstring(strings.Repeat("x", 30)))
}

func TestFooterShowsTheNewKeys(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, "!a/a1"), Run: (&fakeRuns{}).run})
	runCmds(t, m, cmd)

	g.Expect(m.View().Content).To(ContainSubstring("n next   l list   c check all"))
}
