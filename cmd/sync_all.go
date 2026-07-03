package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/sync/v2"
)

var entrySyncAllCmd = &cobra.Command{
	Use:   "sync:all",
	Short: "Sync all entries",
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

		err = sync.All()

		if err != nil {
			fmt.Println("Error syncing entry:", err)
			return
		}

		fmt.Printf("%s\n", s.Render("Successfully synced entries "))
	},
}

func init() {
	rootCmd.AddCommand(entrySyncAllCmd)
}
