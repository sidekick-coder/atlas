package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/sync/v2"
)

var entrySyncCmd = &cobra.Command{
	Use:   "entry:sync",
	Short: "Sync an entry",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		sync := sync.Create(
			app.Drive(),
			app.EntryRepo(),
			app.EntryMetaRepo(),
		)

		s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

		fmt.Printf("%s\n", s.Render("Syncing entry: " + args[0]))

		err = sync.One(args[0])

		if err != nil {
			fmt.Println("Error syncing entry:", err)
			return
		}

		fmt.Printf("%s\n", s.Render("Successfully synced entry: " + args[0]))
	},
}

func init() {
	rootCmd.AddCommand(entrySyncCmd)
}
