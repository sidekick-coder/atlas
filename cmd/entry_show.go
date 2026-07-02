/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
	"charm.land/lipgloss/v2"

)

// entryShowCmd represents the entryShow command
var entryShowCmd = &cobra.Command{
	Use:   "entry:show",
	Short: "Show an entry in the workspace",
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

		entryRepo := entry.New(database)
		entryMetaRepo := entrymeta.New(database)

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

		s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

		fmt.Printf("%s\n", entry.Path)

		for _, m := range metas {
			fmt.Printf("%s: %s\n", s.Render(m.Name), m.Value)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(entryShowCmd)
}
