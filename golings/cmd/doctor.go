package cmd

import (
	"errors"

	"github.com/cleonte/golings/golings/doctor"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func DoctorCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:           "doctor",
		Short:         "Check your Go installation and the exercises workspace",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			checks := doctor.Run(infoFile, doctor.DefaultEnv())
			for _, c := range checks {
				switch c.Status {
				case doctor.OK:
					color.Green("✓ %s: %s", c.Name, c.Detail)
				case doctor.Warn:
					color.Yellow("! %s: %s", c.Name, c.Detail)
				case doctor.Fail:
					color.Red("✗ %s: %s", c.Name, c.Detail)
				}
			}
			if doctor.Failed(checks) {
				return errors.New("doctor found problems")
			}
			return nil
		},
	}
}
