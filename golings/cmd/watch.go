package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

			go WatchEvents(update)

			for {
				go func() {
					for range update {
						fmt.Println("RECEIVER", update)
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
					exs, err := exercises.List(infoFile)
					if err != nil {
						color.Red(err.Error())
						os.Exit(1)
					}
					ui.PrintList(os.Stdout, exs)

				case "hint":
					pathSlice := strings.Split(curFile, "/")
					exIndex := len(pathSlice) - 2

					if exIndex != -1 {
						exs, err := exercises.Find(pathSlice[exIndex], infoFile)
						if err != nil {
							color.Red("Error finding file to hint")
							os.Exit(1)
						}
						color.Yellow(exs.Hint)
					} else {
						color.Red("Error in detect which file to run hint command. Please save the file again and type hint")
					}

				default:
					color.Yellow("only list or hint command are avaliable")
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
		for event := range watcher.Events {
			fmt.Println("EVENT", event)
			if event.Has(fsnotify.Write) {
				fmt.Println("NOTIFY", event.Name)
				updateF <- event.Name
			}
		}
	}()
}

func RunNextExercise(infoFile string) {
	CallClear()
	fmt.Println("RUN NEXT EXEERCISE")
	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
	}

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

func CallClear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			color.Red("Clear terminal command error")
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			color.Red("Clear terminal command error")
		}
	}
}
