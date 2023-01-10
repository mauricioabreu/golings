package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func WatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Run a single exercise",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Create watcher")
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}
			defer watcher.Close()

			path, _ := os.Getwd()
			file_path := fmt.Sprintf("%s/exercises", path)

			err = filepath.WalkDir(file_path, func(path_dir string, d fs.DirEntry, err error) error {
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
						log.Println("event:", event)
						if event.Has(fsnotify.Write) {
							log.Println("modified file:", event.Name)
							cmd := exec.Command("golings run next")
							if err := cmd.Run(); err != nil {
								log.Fatal(err)
							}
						}
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						log.Println("error:", err)
					}
				}
			}()

			// Block main goroutine forever.
			<-make(chan struct{})
		},
	}
}
