package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/tui"
)

// uiCmd represents the ui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the terminal user interface",
	Run: func(cmd *cobra.Command, args []string) {
		err := tui.Run()

		if err != nil {
			fmt.Println("Error launching TUI:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
