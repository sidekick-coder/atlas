/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/sidekick-coder/atlas/internal/store"

)

// entryShowCmd represents the entryShow command
var entryShowCmd = &cobra.Command{
	Use:   "entry:show",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		ws, err := workspace.Get()
		target := args[0]

		if err != nil {
			fmt.Println("Error getting workspace:", err)
			return nil
		}

		dbPath, err := db.Path(ws)

		if err != nil {
			return err
		}

		// start a connection
		conn, err := db.Connect(dbPath)

		if err != nil {
			return err 
		}

		defer conn.Close()

		store := store.New(conn)

		entry, err := store.GetEntryByPath(target)

		if err != nil {
			fmt.Println("Error getting entry:", err)
			return nil
		}

		if entry == nil {
			fmt.Println("Entry not found")
			return nil
		}

		metas, err := store.GetEntryMetasByEntryID(entry.ID)

		if err != nil {
			fmt.Println("Error getting entry metas:", err)
			return nil
		}

		entryType := "file"

		if entry.IsDir {
			entryType = "directory"
		}

		fmt.Printf("%s\n", entry.Path)
		fmt.Printf("\ttype: %s\n", entryType)

		for _, meta := range metas {
			fmt.Printf("\t%s: %s\n", meta.Name, meta.Value)
		}


		return nil

	},
}

func init() {
	rootCmd.AddCommand(entryShowCmd)
}
