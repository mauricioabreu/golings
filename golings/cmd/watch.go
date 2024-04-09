package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func WatchCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Verify exercises when files are edited",
		RunE: func(cmd *cobra.Command, args []string) error {
			RunNextExercise(infoFile)
			reader := bufio.NewReader(os.Stdin)
			update := make(chan string)

			go WatchEvents(update)

			for {
				go func() {
					for range update {
						RunNextExercise(infoFile)
					}
				}()

				cmdString, err := reader.ReadString('\n')
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				cmdStr := strings.TrimSuffix(cmdString, "\n")

				switch cmdStr {
				case "list":
					PrintList(infoFile)
				case "hint":
					PrintHint(infoFile)
				case "quit":
					color.Green("Bye by golings o/")
					os.Exit(0)
				case "exit":
					color.Green("Bye by golings o/")
					os.Exit(0)
				default:
					color.Yellow("only list or hint commands are available")
				}
			}
		},
	}
}

func WatchEvents(updateF chan<- string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	path, _ := os.Getwd()
	directories := fmt.Sprintf("%s/exercises", path)

	err = filepath.WalkDir(directories, func(path_dir string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}
		if d.IsDir() {
			err = watcher.Add(path_dir)

			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error in file path:", err.Error())
	}

	for event := range watcher.Events {
		if event.Has(fsnotify.Write) || event.Has(fsnotify.Rename) {
			updateF <- event.Name
		}
	}
}
