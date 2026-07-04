package cmd

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/actionmanager"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Execute an action for the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0] 
		params := args[1:]
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}

		am, err := actionmanager.New(config)

		if err != nil {
			fmt.Println("Error creating action manager:", err)
			return
		}
		
		err = am.Execute(name, params)

		if err != nil {
			fmt.Println("Error executing action:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(actionCmd)
}
