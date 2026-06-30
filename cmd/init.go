package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize .atlas folder in directory",
	Long: `This command initializes a .atlas folder in the current directory. This folder is used to store metadata about your files and directories.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ws, err := workspace.Get()
		if err != nil {
			return err
		}

		fmt.Println("Workspace:", ws)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
