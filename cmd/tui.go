package cmd

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	tui "github.com/sidekick-coder/atlas/tui/root"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the terminal user interface",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := app.Create()

		if err != nil {
			panic(fmt.Sprintf("Error creating app: %v", err))
		}

		f, err := tea.LogToFile("tui.log", "debug")

		if err != nil {
			fmt.Println("Error creating log file:", err)
			os.Exit(1)
		}

		defer f.Close()

		root := tui.New(a)

		p := tea.NewProgram(root)

		tui.Program = p

		_, err = p.Run()

		if err != nil {
			fmt.Println("Error launching TUI:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
