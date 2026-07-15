package cmd

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan entries",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		drive := app.Drive()

		drive.ScanStream(func(e models.EntryInfo) error {
			fmt.Printf("%s\n", e.Path)

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
