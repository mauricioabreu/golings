package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func WatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Run a single exercise",
		Run: func(cmd *cobra.Command, args []string) {
			color.White("WATCHING ALL")
			// return nil
		},
	}
}
