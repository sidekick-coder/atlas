package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List entries in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.Create()
		showMetas, _ := cmd.Flags().GetStringArray("metas")
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")

		if err != nil {
			fmt.Println(err)
			return nil
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()

		options := entry.ListOptions{
			Query:  args,
			Limit:  limit,
			Offset: offset,
		}

		entries, err := entryRepo.List(options)

		if err != nil {
			fmt.Println("Error listing entries:", err)
			return nil
		}

		for _, entry := range entries {
			fmt.Printf("%s\n", entry.Path)

			if len(showMetas) > 0 {
				metas, err := entryMetaRepo.ListByEntryID(entry.ID)

				if err != nil {
					fmt.Println("Error fetching metadata for entry:", err)
					continue
				}

				utils.PrintMetas(metas)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	rootCmd.PersistentFlags().StringArrayP("metas", "m", []string{}, "Show metadata for the entries")
	rootCmd.PersistentFlags().IntP("limit", "l", 0, "Limit the number of entries to list")
	rootCmd.PersistentFlags().IntP("offset", "o", 0, "Offset the entries to list")
}
