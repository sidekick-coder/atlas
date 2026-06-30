package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/sidekick-coder/atlas/internal/sync"
	"github.com/spf13/cobra"
)

// syncAllCmd represents the syncAll command
var syncAllCmd = &cobra.Command{
	Use:   "sync-all",
	Short: "Sync all files and directories in the workspace",
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

		if err := sync.All(conn, ws); err != nil {
			fmt.Println("Error syncing all files and directories:", err)
			return nil
		}

		fmt.Println("Successfully synced all files and directories in the workspace")

		return nil

	},
}

func init() {
	rootCmd.AddCommand(syncAllCmd)
}
