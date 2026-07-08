package cmd

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/spf13/cobra"
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

		err = app.SetEntryMeta(filename, name, value)

		if err != nil {
			fmt.Println("Error setting metadata:", err)
			return nil
		}

		sync := app.Syncer()

		err = sync.One(args[0])

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
