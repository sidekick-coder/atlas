package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/sync/v2"
)

// setCmd represents the entryShow command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata for an entry in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0] 
		name := args[1]
		value := args[2] 

		app, err := app.Create()

		if err != nil {
			fmt.Println("Error on initialization:", err)
			return nil
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()
		drive := app.Drive()

		entry, err := entryRepo.GetByPath(filename)

		if err != nil {
			fmt.Println("Error showing entry:", err)
			return nil
		}

		info, err := drive.Get(entry.Path)

		if err != nil {
			fmt.Println("Error getting entry info:", err)
			return nil
		}

		handlers := metadata.GetHandlers(info)

		success, err := metadata.Set(info, name, value, handlers)

		if err != nil {
			fmt.Println("Error setting metadata:", err)
			return nil
		}

		if !success {
			fmt.Println("Could not set value:", name)
			return nil
		}

		sync := sync.Create(
			app.Drive(),
			app.EntryRepo(),
			app.EntryMetaRepo(),
		)

		err = sync.One(args[0])

		if err != nil {
			fmt.Println("Error syncing entry:", err)
			return nil
		}

		entry, err = entryRepo.GetByPath(args[0])

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
	rootCmd.AddCommand(setCmd)
}
