package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize .atlas folder in directory",
	Long: `This command initializes a .atlas folder in the current directory. This folder is used to store metadata about your files and directories.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return nil
		}


		dir := config.Get("workspace.atlas_path")

		info, err := os.Stat(dir)

		if err == nil && info.IsDir() {
			fmt.Println(".atlas folder already exists")
			return nil
		}

		fmt.Printf("Creating .atlas folder in %s\n", dir)

		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			return err
		}

		database, err := database.Create(config.Get("workspace.database_path"))

		if err != nil {
			fmt.Println("Error creating database:", err)
			return nil
		}

		err = database.Migrate()

		if err != nil {
			fmt.Println("Error migrating database:", err)
			return nil
		}


		fmt.Println("Initialized .atlas folder and database")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
