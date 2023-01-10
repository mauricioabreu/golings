package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

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
			log.Println("Create watcher")
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}
			defer watcher.Close()

			// Start listening for events.
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

			path, _ := os.Getwd()
			file_path := fmt.Sprintf("%s/exercises", path)

			err = filepath.WalkDir(file_path, func(path_dir string, d fs.DirEntry, err error) error {
				if err != nil {
					log.Fatal(err)
					return err
				}
				if d.IsDir() {
					log.Printf("Added %s to watch\n", path_dir)
					err = watcher.Add(file_path)

					if err != nil {
						log.Fatal(err)
					}
				}
				return nil
			})

			if err != nil {
				log.Fatal("Error in file path:", err.Error())
			}

			// Block main goroutine forever.
			<-make(chan struct{})
		},
	}
}
