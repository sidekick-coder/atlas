/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/app"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all entries and metadata",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		entryRepo := app.EntryRepo()
		entryMetaRepo := app.EntryMetaRepo()

		entryMetaRepo.DeleteAll()
		entryRepo.DeleteAll()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
