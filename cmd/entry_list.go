package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/spf13/cobra"
	"charm.land/lipgloss/v2"
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

		s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))


		for _, entry := range entries {
			fmt.Printf("%s\n", entry.Path)

			meta, err := entryMetaRepo.ListByEntryID(entry.ID)

			if err != nil {
				fmt.Println("Error fetching metadata for entry:", err)
				continue
			}

			for _, m := range meta {
				fmt.Printf("  %s: %s\n", s.Render(m.Name), m.Value)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
