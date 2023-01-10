package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func WatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Run a single exercise",
		Run: func(cmd *cobra.Command, args []string) {
			color.White("WATCHING ALL")
			// return nil
			// watcher("../../exercises")
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}

			defer watcher.Close()

			go func() {
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						log.Println("event:", event)
						if event.Has(fsnotify.Write) {
							log.Println("modified file:", event.Name)
						}
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						log.Println("error:", err)
					}
				}
			}()

			path, err := os.Getwd()
			if err := watcher.Add(path); err != nil {
				log.Fatal("failed to watch %s, error: %s", path, err.Error())
			}
			fmt.Println("PATH:", path)
			if err != nil {
				log.Fatal(err)
			}
			/* ... do stuff ... */

			if err != nil {
				log.Fatal(err)
			}

			// Block main goroutine forever.
			<-make(chan struct{})

		},
	}
}
