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
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/mauricioabreu/golings/golings/ui"
	"github.com/spf13/cobra"
)

func WatchCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Run a single exercise",
		RunE: func(cmd *cobra.Command, args []string) error {
			RunNextExercise(infoFile)
			reader := bufio.NewReader(os.Stdin)
			update := make(chan string)
			var curFile string

			for {
				go WatchEvents(update)

				go func() {
					for f := range update {
						// @TODO: use this filename to command hint
						fmt.Println("FILE UPDATED:", f)
						RunNextExercise(infoFile)
						curFile = <-update
					}
				}()

				cmdString, err := reader.ReadString('\n')
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

				cmdStr := strings.TrimSuffix(cmdString, "\n")

				switch cmdStr {
				case "list":
					log.Println("List command", cmdString)
					exs, err := exercises.List(infoFile)
					if err != nil {
						color.Red(err.Error())
						os.Exit(1)
					}
					ui.PrintList(os.Stdout, exs)

				case "hint":
					pathSlice := strings.Split(curFile, "/")
					exIndex := len(pathSlice) - 2
					exs, err := exercises.Find(pathSlice[exIndex], infoFile)

					if err != nil {
						color.Red(err.Error())
						os.Exit(1)
					}
					color.Yellow(exs.Hint)
				default:
					color.Yellow("only list or hint command are avaliable")
				}
			}
			return nil
		},
	}
}

func WatchEvents(updateF chan<- string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	path, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/exercises", path)

	err = filepath.WalkDir(filePath, func(path_dir string, d fs.DirEntry, err error) error {
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

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					updateF <- event.Name
				}
			}
		}
	}()
}

func RunNextExercise(infoFile string) {
	exercise, err := exercises.NextPending(infoFile)
	result, err := exercise.Run()

	if err != nil {
		color.Cyan("Failed to compile the exercise %s\n\n", result.Exercise.Path)
		color.White("Check the output below: \n\n")
		color.Red(result.Err)
		color.Red(result.Out)
		color.Yellow("If you feel stuck, ask a hint by executing `golings hint %s`", result.Exercise.Name)
	} else {
		color.Green("Congratulations!\n\n")
		color.Green("Here is the output of your program:\n\n")
		color.Cyan(result.Out)
		if result.Exercise.State() == exercises.Pending {
			color.White("Remove the 'I AM NOT DONE' from the file to keep going\n")
			color.Red("exercise is still pending")
		}
	}
}
