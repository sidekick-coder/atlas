package cmd

import (
	"fmt"
	"log/slog"
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

		file, err := os.OpenFile("tui.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println("Error opening log file:", err)
			os.Exit(1)
		}

		defer file.Close()

		logger := slog.New(slog.NewJSONHandler(file, nil))

		slog.SetDefault(logger)

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
