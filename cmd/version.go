package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version string = "0.8.0"

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Displays the current version",

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(os.Stdout, "Current version: %s", version)

			return nil
		},
	}

	return cmd
}
