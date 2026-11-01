package cmd

import (
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cleonte/golings/golings/tui"
	"github.com/cleonte/golings/golings/watcher"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// debounce is how long the watcher waits for a burst of file events to settle.
const debounce = 150 * time.Millisecond

func WatchCmd(infoFile string) *cobra.Command {
	var plain bool

	watchCmd := &cobra.Command{
		Use:   "watch",
		Short: "Verify exercises when files are edited",
		Long: `Runs the next exercise whenever you save a file.

In a terminal this opens a full-screen view. Use --plain for simple line-by-line
output; that is also what you get when input or output is not a terminal, for
example when piping.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			root := filepath.Join(filepath.Dir(infoFile), "exercises")
			changes, err := watcher.Watch(ctx, root, debounce)
			if err != nil {
				color.Red("Can't watch %s: %s", root, err)
				color.Yellow("Run golings from the folder that holds the exercises, create one with `golings init`, or try `golings doctor`.")
				return err
			}

			if !plain && isTerminal(cmd.InOrStdin()) && isTerminal(cmd.OutOrStdout()) {
				return tui.Run(ctx, tui.Config{InfoFile: infoFile, Changes: changes})
			}

			watchLoop(cmd.InOrStdin(), changes, watchActions{
				run:  func() { RunNextExercise(infoFile) },
				list: func() { PrintList(infoFile) },
				hint: func() { PrintHint(infoFile) },
			})
			return nil
		},
	}
	watchCmd.Flags().BoolVar(&plain, "plain", false, "print plain lines instead of the full-screen view")

	return watchCmd
}

// isTerminal reports whether f is an interactive terminal (and not, say, a pipe).
func isTerminal(f any) bool {
	file, ok := f.(*os.File)
	return ok && term.IsTerminal(int(file.Fd()))
}

type watchActions struct {
	run, list, hint func()
}

// watchLoop runs the next exercise now and whenever files change, and answers
// the commands typed by the learner, until they quit or input ends. Everything
// happens in this one loop, so a run never overlaps with another run or with a
// command.
func watchLoop(in io.Reader, changes <-chan []string, actions watchActions) {
	actions.run()

	lines := make(chan string)
	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			lines <- strings.TrimSpace(scanner.Text())
		}
	}()

	for {
		select {
		case _, ok := <-changes:
			if !ok {
				return
			}
			actions.run()

		case line, ok := <-lines:
			if !ok {
				return
			}
			switch line {
			case "list":
				actions.list()
			case "hint":
				actions.hint()
			case "quit", "exit":
				color.Green("Bye by golings o/")
				return
			default:
				color.Yellow("only list or hint commands are available")
			}
		}
	}
}
