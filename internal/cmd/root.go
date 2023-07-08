package cmd

import (
	"fmt"
	"os"

	"github.com/ChrisWiegman/gh-backup/internal/backup"
	"github.com/spf13/cobra"
)

var flagVersion bool

func Execute() {
	// Setup the cobra command
	cmd := &cobra.Command{
		Use:   "gh-backup <GitHub Username>",
		Short: "GH Backup is a simple script to backup all public GitHub repos for a specified account.",
		Args:  cobra.MaximumNArgs(1),
		Run:   backup.ExecuteBackup,
	}

	cmd.PersistentFlags().BoolVarP(&flagVersion, "version", "v", false, "Display version information for the gh-backup client.")

	// Execute anything we need to
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
