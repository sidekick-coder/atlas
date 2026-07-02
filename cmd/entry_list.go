package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "entry:list",
	Short: "List entries in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return nil
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()

		entries, err := entryRepo.List()

		if err != nil {
			fmt.Println("Error listing entries:", err)
			return nil
		}

		s := lipgloss.
			NewStyle().
			Bold(true).
			Foreground(lipgloss.Red)

		for _, entry := range entries {
			fmt.Printf("%s\n", s.Render(entry.Path))

			metas, err := entryMetaRepo.ListByEntryID(entry.ID)

			if err != nil {
				fmt.Println("Error fetching metadata for entry:", err)
				continue
			}

			utils.PrintMetas(metas)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
