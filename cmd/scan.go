package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/sidekick-coder/atlas/internal/drive"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the workspace for files and directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		ws, err := workspace.Get()
		if err != nil {
			return err
		}

		entries, err := drive.Scan(ws)

		if err != nil {
			return err
		}

		for _, entry := range entries {
			fmt.Printf("%s (dir: %t)\n", entry.Path, entry.IsDir)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
