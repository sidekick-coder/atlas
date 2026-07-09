package cmd

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/spf13/cobra"
)

var entryExtractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract an entry from the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]

		app, err := app.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}

		entry, err := app.Drive().Get(filepath)

		if err != nil {
			fmt.Println("Error getting entry:", err)
			return
		}

		h, err := metadata.Handler(entry)

		if err != nil {
			fmt.Println("Error creating metadata:", err)
			return
		}

		metas, err := h.Extract()

		if err != nil {
			fmt.Println("Error extracting metadata:", err)
			return
		}

		fmt.Printf("%s\n", entry.Path)

		utils.PrintMetas(metas)

	},
}

func init() {
	rootCmd.AddCommand(entryExtractCmd)
}
