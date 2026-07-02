package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/utils"

)

// entryShowCmd represents the entryShow command
var entryShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show an entry in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.Create()

		if err != nil {
			fmt.Println("Error on initialization:", err)
			return nil
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()

		entry, err := entryRepo.GetByPath(args[0])

		if err != nil {
			fmt.Println("Error showing entry:", err)
			return nil
		}

		metas, err := entryMetaRepo.ListByEntryID(entry.ID)

		if err != nil {
			fmt.Println("Error listing entry metas:", err)
			return nil
		}

		fmt.Printf("%s\n", entry.Path)

		utils.PrintMetas(metas)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(entryShowCmd)
}
