package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)

var entryExtractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract an entry from the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}

		drive, err := drive.New(config.Get("workspace.path"))

		if err != nil {
			fmt.Println("Error creating drive:", err)
			return
		}

		entry, err := drive.Get(filepath)

		if err != nil {
			fmt.Println("Error getting entry:", err)
			return
		}

		meta, err := metadata.Create(entry)

		if err != nil {
			fmt.Println("Error creating metadata:", err)
			return
		}

		data, err := meta.Extract()

		if err != nil {
			fmt.Println("Error extracting metadata:", err)
			return
		}

		fmt.Printf("%s\n", entry.Path)

		utils.PrintMetas(data)

	},
}

func init() {
	rootCmd.AddCommand(entryExtractCmd)
}
