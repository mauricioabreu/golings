package tui

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	tea "charm.land/bubbletea/v2"
	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/gomega"
)

const original = "// I AM NOT DONE\npackage main\n\nfunc main() {}\n"

func press(m *Model, keys ...string) tea.Cmd {
	var cmd tea.Cmd
	for _, key := range keys {
		code, _ := firstRune(key)
		msg := tea.KeyPressMsg{Code: code, Text: key}
		switch key {
		case "esc":
			msg = tea.KeyPressMsg{Code: tea.KeyEscape}
		case "down":
			msg = tea.KeyPressMsg{Code: tea.KeyDown}
		case "up":
			msg = tea.KeyPressMsg{Code: tea.KeyUp}
		}
		_, cmd = m.Update(msg)
	}
	return cmd
}

func firstRune(s string) (rune, bool) {
	for _, r := range s {
		return r, true
	}
	return 0, false
}

// withExercise starts a model on a single unsolved exercise "a1" whose run has
// finished, with a hint, a solution and the original exercise to reset from.
func withExercise(t *testing.T, mutate ...func(*Config)) (*Model, string) {
	t.Helper()
	noColor(t)
	info := newWorkspace(t, "!a/a1")
	root := filepath.Dir(info)
	manifest, err := os.ReadFile(info)
	if err != nil {
		t.Fatal(err)
	}
	// the same manifest, with the paths the embedded copy would have
	embeddedManifest := strings.ReplaceAll(string(manifest), filepath.Join(root, "exercises"), "exercises")
	embeddedManifest = strings.Replace(embeddedManifest, `hint = ""`, `hint = "Think about the return value.\nThen about the arguments."`, 1)
	if err := os.WriteFile(info, []byte(embeddedManifest), 0o644); err != nil {
		t.Fatal(err)
	}
	// the exercise paths are relative to the info file's folder
	t.Chdir(root)

	cfg := Config{
		InfoFile: info,
		Run:      (&fakeRuns{}).run,
		Source: fstest.MapFS{
			"info.toml":              {Data: []byte(embeddedManifest)},
			"exercises/a/a1/main.go": {Data: []byte(original)},
		},
		Root:      root,
		Solutions: fstest.MapFS{"solutions/a/a1/main.go": {Data: []byte("package main\n\nfunc main() {\n\tprintln(\"solved\")\n}\n")}},
	}
	for _, f := range mutate {
		f(&cfg)
	}
	m, cmd := start(t, cfg)
	runCmds(t, m, cmd)
	return m, filepath.Join(root, "exercises", "a", "a1", "main.go")
}

func TestHintOpensAndCloses(t *testing.T) {
	g := NewWithT(t)
	m, _ := withExercise(t)

	press(m, "h")
	g.Expect(m.View().Content).To(ContainSubstring("Hint for a1"))
	g.Expect(m.View().Content).To(ContainSubstring("Think about the return value."))
	g.Expect(m.View().Content).To(ContainSubstring("Then about the arguments."))
	g.Expect(m.View().Content).To(ContainSubstring("esc back"))

	press(m, "h")
	g.Expect(m.View().Content).NotTo(ContainSubstring("Hint for a1"))

	press(m, "h", "esc")
	g.Expect(m.View().Content).NotTo(ContainSubstring("Hint for a1"))
	g.Expect(m.View().Content).To(ContainSubstring("h hint"))
}

func TestHintWithoutOne(t *testing.T) {
	g := NewWithT(t)
	m, _ := withExercise(t)
	m.exercise.Hint = ""

	press(m, "h")

	g.Expect(m.View().Content).To(ContainSubstring("no hint for this exercise"))
}

func TestSolutionIsHiddenUntilConfirmed(t *testing.T) {
	g := NewWithT(t)
	m, _ := withExercise(t)

	press(m, "s")
	g.Expect(m.View().Content).To(ContainSubstring("Show the solution of a1?"))
	g.Expect(m.View().Content).NotTo(ContainSubstring("solved"))

	press(m, "n")
	g.Expect(m.View().Content).NotTo(ContainSubstring("solved"))
	g.Expect(m.mode).To(Equal(modeResult))

	press(m, "s", "x")
	g.Expect(m.View().Content).NotTo(ContainSubstring("solved"), "only y says yes")

	press(m, "s", "y")
	g.Expect(m.View().Content).To(ContainSubstring("Solution for a1"))
	g.Expect(m.View().Content).To(ContainSubstring(`println("solved")`))
	g.Expect(m.View().Content).NotTo(ContainSubstring("\t"), "tabs are expanded")

	press(m, "esc")
	g.Expect(m.View().Content).NotTo(ContainSubstring("solved"))
}

func TestSolutionNotAvailable(t *testing.T) {
	g := NewWithT(t)

	none, _ := withExercise(t, func(c *Config) { c.Solutions = nil })
	press(none, "s")
	g.Expect(none.View().Content).To(ContainSubstring("Solutions are not available here."))
	g.Expect(none.mode).To(Equal(modeResult))

	missing, _ := withExercise(t, func(c *Config) { c.Solutions = fstest.MapFS{} })
	press(missing, "s", "y")
	g.Expect(missing.View().Content).To(ContainSubstring("No solution available for a1 yet."))
	g.Expect(missing.mode).To(Equal(modeResult))
}

func TestResetNeedsConfirmationAndKeepsWorkOtherwise(t *testing.T) {
	g := NewWithT(t)
	m, file := withExercise(t)
	g.Expect(os.WriteFile(file, []byte("package main // my work\n"), 0o644)).To(Succeed())

	press(m, "r")
	g.Expect(m.View().Content).To(ContainSubstring("Reset a1 to its original state?"))
	g.Expect(m.View().Content).To(ContainSubstring("discards your changes"))

	for _, answer := range []string{"n", "esc", "x"} {
		press(m, "r", answer)
		data, err := os.ReadFile(file)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(data)).To(Equal("package main // my work\n"), "answer %q", answer)
	}
}

func TestResetRestoresTheOriginalAndRunsAgain(t *testing.T) {
	g := NewWithT(t)
	fake := &fakeRuns{}
	m, file := withExercise(t, func(c *Config) { c.Run = fake.run })
	g.Expect(os.WriteFile(file, []byte("package main // my work\n"), 0o644)).To(Succeed())
	runsBefore := len(fake.started)
	idBefore := m.runID

	cmd := press(m, "r", "y")

	data, err := os.ReadFile(file)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(data)).To(Equal(original))
	g.Expect(cmd).NotTo(BeNil(), "the restored exercise is run again")
	g.Expect(m.runID).To(BeNumerically(">", idBefore))
	g.Expect(m.status).To(Equal(statusRunning))
	g.Expect(m.View().Content).To(ContainSubstring("Reset exercises/a/a1/main.go to its original state."))
	runCmds(t, m, cmd)
	g.Expect(len(fake.started)).To(Equal(runsBefore + 1))
}

func TestResetNotAvailableWithoutASource(t *testing.T) {
	g := NewWithT(t)
	m, _ := withExercise(t, func(c *Config) { c.Source = nil })

	press(m, "r")

	g.Expect(m.View().Content).To(ContainSubstring("Reset is not available here."))
	g.Expect(m.mode).To(Equal(modeResult))
}

func TestResetFailureIsReported(t *testing.T) {
	g := NewWithT(t)
	m, file := withExercise(t, func(c *Config) { c.Source = fstest.MapFS{} })
	g.Expect(os.WriteFile(file, []byte("package main // my work\n"), 0o644)).To(Succeed())

	press(m, "r", "y")

	g.Expect(m.View().Content).To(ContainSubstring("Reset failed:"))
	data, _ := os.ReadFile(file)
	g.Expect(string(data)).To(Equal("package main // my work\n"))
}

func TestQuitWorksInEveryMode(t *testing.T) {
	for _, keys := range [][]string{{"h"}, {"s"}, {"s", "y"}, {"r"}} {
		g := NewWithT(t)
		m, _ := withExercise(t)
		press(m, keys...)

		cmd := press(m, "q")

		g.Expect(exec(t, cmd)).To(BeAssignableToTypeOf(tea.QuitMsg{}), "after %v", keys)
	}
}

func TestKeysDoNothingWithoutAnExercise(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	m, _ := start(t, Config{InfoFile: newWorkspace(t, "a/a1"), Run: (&fakeRuns{}).run, Solutions: fstest.MapFS{}, Source: fstest.MapFS{}})
	g.Expect(m.status).To(Equal(statusAllDone))

	press(m, "h", "s", "r")

	g.Expect(m.mode).To(Equal(modeResult))
}

func TestScrolling(t *testing.T) {
	g := NewWithT(t)
	m, _ := withExercise(t, func(c *Config) {
		c.Run = (&fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) {
			return exercises.Result{Exercise: ex, Out: strings.Repeat("filler\n", 40) + "THE LAST LINE\n"}, nil
		}}).run
	})
	m.Update(tea.WindowSizeMsg{Width: 60, Height: 12})
	g.Expect(m.View().Content).NotTo(ContainSubstring("THE LAST LINE"))

	for range 100 {
		press(m, "down")
	}
	g.Expect(m.scroll).To(Equal(m.maxScroll()), "cannot scroll past the end")
	g.Expect(m.View().Content).To(ContainSubstring("THE LAST LINE"))

	for range 100 {
		press(m, "up")
	}
	g.Expect(m.scroll).To(Equal(0))
	g.Expect(m.View().Content).NotTo(ContainSubstring("THE LAST LINE"))
}

func TestANewRunScrollsBackToTheTopAndClosesPanels(t *testing.T) {
	g := NewWithT(t)
	m, _ := withExercise(t)
	press(m, "h")
	g.Expect(m.mode).To(Equal(modeHint))

	m.Update(changesMsg{files: []string{"x.go"}})

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(m.scroll).To(Equal(0))
}
