package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/sync/v2"
)

var entrySyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync an entry",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0] 

		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()

		sync := sync.Create(
			app.Drive(),
			app.EntryRepo(),
			app.EntryMetaRepo(),
		)

		s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

		fmt.Printf("%s\n", s.Render("Syncing entry: " + filepath))

		err = sync.One(filepath)

		if err != nil {
			fmt.Println("Error syncing entry:", err)
			return
		}

		entry, err := entryRepo.GetByPath(filepath)

		if err != nil {
			fmt.Println("Error showing entry:", err)
			return 
		}

		metas, err := entryMetaRepo.ListByEntryID(entry.ID)

		if err != nil {
			fmt.Println("Error listing entry metas:", err)
			return
		}

		fmt.Printf("%s\n", entry.Path)

		utils.PrintMetas(metas)

	},
}

func init() {
	rootCmd.AddCommand(entrySyncCmd)
}
