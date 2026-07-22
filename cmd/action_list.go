package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/spf13/cobra"
)

var actionListCmd = &cobra.Command{
	Use:   "action:list",
	Short: "List actions available",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.Create()

		if err != nil {
			fmt.Println("Error creating app:", err)
			return
		}

		actions, err := app.Action.List()

		if err != nil {
			fmt.Println("Error listing actions:", err)
			return
		}

		if len(actions) == 0 {
			fmt.Println("No actions found.")
			return
		}

		for _, action := range actions {
			fmt.Printf("%s\n", action.ID)
			fmt.Printf("- ID: %s\n", action.ID)
			fmt.Printf("- Type: %s\n", action.Type)
			fmt.Printf("- Options: %v\n", action.Options)
		}

	},
}

func init() {
	rootCmd.AddCommand(actionListCmd)
}
