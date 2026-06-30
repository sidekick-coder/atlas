package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize .atlas folder in directory",
	Long: `This command initializes a .atlas folder in the current directory. This folder is used to store metadata about your files and directories.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ws, err := workspace.Get()
		if err != nil {
			return err
		}

		dir := fmt.Sprintf("%s/.atlas", ws) 

		info, err := os.Stat(dir)

		if err == nil && info.IsDir() {
			fmt.Println(".atlas folder already exists")
			return nil
		}

		fmt.Printf("Creating .atlas folder in %s\n", ws)

		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			return err
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

		// run migrations
		if err := db.Migrate(conn); err != nil {
			return err
		}

		fmt.Println("Initialized .atlas folder and database")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
