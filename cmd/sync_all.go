package cmd

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/syncer"
	"github.com/spf13/cobra"
)

var syncAllCmd = &cobra.Command{
	Use:   "sync:all",
	Short: "Sync all entries",
	Run: func(cmd *cobra.Command, args []string) {
		concurrency, err := cmd.Flags().GetInt("concurrency")

		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		green := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
		red := lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

		onSuccess := func(path string) {
			fmt.Printf("%s\n", green.Render(path))
		}

		onError := func(path string, err error) {
			fmt.Printf("%s\n", red.Render(path))
			fmt.Printf("Error: %v\n", err)
		}

		onComplete := func(result syncer.Result) {
			fmt.Printf("\n")
			fmt.Printf("Scanned: %d\n", result.Scanned)
			fmt.Printf("Extracted: %d\n", result.Extracted)
			fmt.Printf("Written: %d\n", result.Written)
			fmt.Printf("Time: %.3f s\n", result.Time.Seconds())
			fmt.Printf("Concurrency: %d\n", concurrency)
		}

		syncer.Create().
			SetConfig(app.Config()).
			SetDatabase(app.Database()).
			SetDrive(app.Drive()).
			SetConcurrency(concurrency).
			OnSuccess(onSuccess).
			OnError(onError).
			OnComplete(onComplete).
			All()
	},
}

func init() {
	rootCmd.AddCommand(syncAllCmd)
	syncAllCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent workers")
}
