package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List entries in the workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		ws, err := workspace.Get()

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

		entries, err := db.SelectEntries(conn)

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
