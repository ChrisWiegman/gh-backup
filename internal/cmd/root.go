package cmd

import (
	"fmt"
	"os"

	"github.com/ChrisWiegman/gh-backup/internal/backup"
	"github.com/spf13/cobra"
)

func Execute() {
	// Setup the cobra command
	cmd := &cobra.Command{
		Use:   "gh-backup <GitHub Username>",
		Short: "GH Backup is a simple script to backup all public GitHub repos for a specified account.",
		Args:  cobra.ExactArgs(1),
		Run:   backup.ExecuteBackup,
	}

	// Execute anything we need to
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
