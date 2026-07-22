package cmd

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/tui/features/theme"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List entries in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		showMetas, _ := cmd.Flags().GetBool("metas")
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")
		query, _ := cmd.Flags().GetStringSlice("query")

		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return nil
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()

		options := entry.ListOptions{
			Query:  []string{},
			Limit:  limit,
			Offset: offset,
		}

		if len(query) > 0 {
			options.Query = query
		}

		entries, err := entryRepo.List(options)

		if err != nil {
			fmt.Println("Error listing entries:", err)
			return nil
		}

		ks := lipgloss.NewStyle().Width(20).Bold(true).Foreground(lipgloss.Color(theme.Current.Primary))
		vs := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Current.Secondary))

		for _, entry := range entries {
			fmt.Printf("%s\n", entry.Path)

			if showMetas {
				metas, err := entryMetaRepo.ListByEntryID(entry.ID)

				if err != nil {
					fmt.Println("Error fetching metadata for entry:", err)
					continue
				}

				for _, m := range metas {
					v := m.Value

					v = strings.ReplaceAll(v, "\n", "\\n")

					if len(v) > 60 {
						v = v[:60] + "..."
					}

					fmt.Printf("  %s: %s\n", ks.Render(m.Name), vs.Render(v))
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	rootCmd.PersistentFlags().StringSliceP("query", "q", []string{}, "Query to filter entries")
	rootCmd.PersistentFlags().BoolP("metas", "m", false, "Show metadata for each entry")
	rootCmd.PersistentFlags().IntP("limit", "l", 0, "Limit the number of entries to list")
	rootCmd.PersistentFlags().IntP("offset", "o", 0, "Offset the entries to list")
}
