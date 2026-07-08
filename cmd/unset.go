package cmd

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/spf13/cobra"
)

// unsetCmd represents the entryShow command
var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Set metadata for an entry in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		name := args[1]

		app, err := app.Create()

		if err != nil {
			fmt.Println("Error on initialization:", err)
			return nil
		}

		err = app.UnsetEntryMeta(filename, name)

		if err != nil {
			fmt.Println("Error setting metadata:", err)
			return nil
		}

		sync := app.Syncer()

		err = sync.One(args[0])

		if err != nil {
			fmt.Println("Error syncing entry:", err)
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unsetCmd)
}
