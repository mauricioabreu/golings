package tui

import (
	"os"
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	. "github.com/onsi/gomega"
)

// feed runs a command, and the commands that come out of feeding its message to
// the model, like the Bubble Tea runtime would. A command that blocks (a
// cursor blinking, the wait for a save) is left alone.
func feed(t *testing.T, m *Model, cmd tea.Cmd) {
	t.Helper()
	if cmd == nil {
		return
	}
	done := make(chan tea.Msg, 1)
	go func() { done <- cmd() }()
	var got tea.Msg
	select {
	case got = <-done:
	case <-time.After(50 * time.Millisecond):
	}
	switch msg := got.(type) {
	case nil:
	case tea.BatchMsg:
		for _, c := range msg {
			feed(t, m, c)
		}
	default:
		_, next := m.Update(msg)
		feed(t, m, next)
	}
}

// typeKeys presses keys one by one, feeding back what the model asks for.
func typeKeys(t *testing.T, m *Model, keys ...string) {
	t.Helper()
	for _, key := range keys {
		feed(t, m, press(m, key))
	}
}

// listed starts a model on the exercises, with the first run finished.
func listed(t *testing.T, fake *fakeRuns, specs ...string) *Model {
	t.Helper()
	noColor(t)
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, specs...), Run: fake.run})
	runCmds(t, m, cmd)
	m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	return m
}

func TestLShowsEveryExerciseWithItsState(t *testing.T) {
	g := NewWithT(t)
	m := listed(t, &fakeRuns{}, "a/a1", "!a/a2", "!b/b1")

	typeKeys(t, m, "l")
	view := plain(m.View().Content)

	g.Expect(m.mode).To(Equal(modeList))
	g.Expect(view).To(MatchRegexp(`✔ a1\s+a`))
	g.Expect(view).To(MatchRegexp(`> · a2\s+a`), "the cursor is on the current exercise")
	g.Expect(view).To(MatchRegexp(`· b1\s+b`))
	g.Expect(view).To(ContainSubstring("enter open   / filter"))
}

func TestEnterOpensTheExerciseUnderTheCursor(t *testing.T) {
	g := NewWithT(t)
	fake := &fakeRuns{}
	m := listed(t, fake, "a/a1", "!a/a2", "!b/b1")

	typeKeys(t, m, "l", "down", "enter")

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(m.exercise.Name).To(Equal("b1"))
	g.Expect(fake.started).To(ContainElement("b1"))
	g.Expect(plain(m.View().Content)).To(ContainSubstring("b1"))
}

func TestAnExerciseFromTheListStaysUntilNMovesOn(t *testing.T) {
	g := NewWithT(t)
	fake := &fakeRuns{}
	m := listed(t, fake, "!a/a1", "!a/a2", "!b/b1")

	typeKeys(t, m, "l", "down", "down", "enter")
	g.Expect(m.exercise.Name).To(Equal("b1"))

	// A save runs the same exercise again, not the first pending one.
	_, cmd := m.Update(changesMsg{files: []string{"b1"}})
	feed(t, m, cmd)
	g.Expect(m.exercise.Name).To(Equal("b1"))

	// n takes out its marker and goes back to the first pending one.
	typeKeys(t, m, "n")
	g.Expect(marker(t, m.exercise.Path)).To(BeTrue(), "now running a1")
	g.Expect(m.exercise.Name).To(Equal("a1"))
	g.Expect(m.pinned).To(BeEmpty())
}

func TestAnExerciseCanBeOpenedWhenEverythingIsDone(t *testing.T) {
	g := NewWithT(t)
	fake := &fakeRuns{}
	m := listed(t, fake, "a/a1", "a/a2")
	g.Expect(m.status).To(Equal(statusAllDone))

	typeKeys(t, m, "l", "down", "enter")

	g.Expect(m.status).To(Equal(statusFinished))
	g.Expect(m.exercise.Name).To(Equal("a2"))
	g.Expect(fake.started).To(ContainElement("a2"))
}

func TestEscLeavesTheListWithoutChangingAnything(t *testing.T) {
	g := NewWithT(t)
	fake := &fakeRuns{}
	m := listed(t, fake, "!a/a1", "!a/a2")
	started := len(fake.started)

	typeKeys(t, m, "l", "down", "esc")

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(m.exercise.Name).To(Equal("a1"))
	g.Expect(fake.started).To(HaveLen(started))
}

func TestTheListCanBeFilteredByNameOrTopic(t *testing.T) {
	g := NewWithT(t)
	m := listed(t, &fakeRuns{}, "!a/a1", "!a/a2", "!maps/m1")

	typeKeys(t, m, "l", "/", "m", "a", "p", "s")
	g.Expect(plain(m.View().Content)).To(ContainSubstring("enter apply filter"))
	typeKeys(t, m, "enter")
	view := plain(m.View().Content)

	g.Expect(view).To(ContainSubstring("m1"))
	g.Expect(view).NotTo(ContainSubstring("a2"))

	typeKeys(t, m, "enter")
	g.Expect(m.exercise.Name).To(Equal("m1"))
}

func TestQIsALetterWhileFilteringAndQuitsOtherwise(t *testing.T) {
	g := NewWithT(t)
	m := listed(t, &fakeRuns{}, "!a/a1", "!q/quiet")

	typeKeys(t, m, "l", "/", "q")
	g.Expect(m.mode).To(Equal(modeList))
	g.Expect(m.list.FilterValue()).To(Equal("q"))

	typeKeys(t, m, "esc")
	g.Expect(m.mode).To(Equal(modeList), "esc only cancelled the filter")
	g.Expect(m.list.SettingFilter()).To(BeFalse())

	cmd := press(m, "q")
	g.Expect(exec(t, cmd)).To(BeAssignableToTypeOf(tea.QuitMsg{}))
}

func TestEscClearsAFilterBeforeLeaving(t *testing.T) {
	g := NewWithT(t)
	m := listed(t, &fakeRuns{}, "!a/a1", "!b/b1")

	typeKeys(t, m, "l", "/", "b", "enter")
	g.Expect(m.list.IsFiltered()).To(BeTrue())

	typeKeys(t, m, "esc")
	g.Expect(m.mode).To(Equal(modeList))
	g.Expect(m.list.IsFiltered()).To(BeFalse())

	typeKeys(t, m, "esc")
	g.Expect(m.mode).To(Equal(modeResult))
}

func TestAFilterWithoutMatchesIsClearedByEnter(t *testing.T) {
	g := NewWithT(t)
	fake := &fakeRuns{}
	m := listed(t, fake, "!a/a1", "!b/b1")
	started := len(fake.started)

	typeKeys(t, m, "l", "/", "z", "z", "z", "enter")

	g.Expect(m.mode).To(Equal(modeList))
	g.Expect(m.list.IsFiltered()).To(BeFalse())
	g.Expect(fake.started).To(HaveLen(started))
}

func TestTheListShowsExercisesThatNeedANewerGo(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := newWorkspace(t, "!a/a1", "!a/a2")
	data, err := os.ReadFile(info)
	if err != nil {
		t.Fatal(err)
	}
	data = []byte(strings.Replace(string(data), "name = \"a2\"\n", "name = \"a2\"\ngo_version = \"99.0\"\n", 1))
	if err := os.WriteFile(info, data, 0o644); err != nil {
		t.Fatal(err)
	}
	m, cmd := start(t, Config{InfoFile: info, Run: (&fakeRuns{}).run})
	runCmds(t, m, cmd)

	typeKeys(t, m, "l")

	g.Expect(plain(m.View().Content)).To(MatchRegexp(`• a2\s+needs Go 99.0`))
}

func TestTheListFollowsTheSizeOfTheScreen(t *testing.T) {
	g := NewWithT(t)
	m := listed(t, &fakeRuns{}, "!a/a1", "!a/a2", "!a/a3", "!a/a4", "!a/a5", "!a/a6", "!a/a7", "!a/a8")

	typeKeys(t, m, "l")
	m.Update(tea.WindowSizeMsg{Width: 40, Height: 8})

	g.Expect(m.list.Height()).To(Equal(m.available()))
	g.Expect(strings.Count(m.View().Content, "\n")).To(BeNumerically("<=", 8))
}

func TestLWithoutExercisesSaysSo(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	info := newWorkspace(t)
	m, cmd := start(t, Config{InfoFile: info, Run: (&fakeRuns{}).run})
	runCmds(t, m, cmd)

	typeKeys(t, m, "l")

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(plain(m.View().Content)).To(ContainSubstring("There are no exercises to list."))
}
