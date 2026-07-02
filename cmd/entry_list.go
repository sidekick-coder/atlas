package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "entry:list",
	Short: "List entries in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return nil
		}

		database, err := database.Create(config.Get("workspace.database_path"))

		if err != nil {
			fmt.Println("Error creating database:", err)
			return nil
		}

		repo := entry.New(database)

		entries, err := repo.List()

		if err != nil {
			fmt.Println("Error listing entries:", err)
			return nil
		}

		for _, entry := range entries {
			fmt.Printf("%s (dir: %t)\n", entry.Path, entry.IsDir)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
