package cmd

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/syncer"
	"github.com/spf13/cobra"
)

var syncAllCmd = &cobra.Command{
	Use:   "sync:all",
	Short: "Sync all entries",
	Run: func(cmd *cobra.Command, args []string) {
		concurrency := 1
		batchSize := 10
		detail := false

		if cmd.Flags().Changed("concurrency") {
			concurrency, _ = cmd.Flags().GetInt("concurrency")
		}

		if cmd.Flags().Changed("batch-size") {
			batchSize, _ = cmd.Flags().GetInt("batch-size")
		}

		if cmd.Flags().Changed("detail") {
			detail, _ = cmd.Flags().GetBool("detail")
		}

		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		green := lipgloss.NewStyle().Foreground(lipgloss.Green)
		red := lipgloss.NewStyle().Foreground(lipgloss.Red)

		onComplete := func(result syncer.Result) {
			fmt.Printf("\n")
			fmt.Printf("Batch Size: %d\n", result.BatchSize)
			fmt.Printf("Concurrency: %d\n", concurrency)
			fmt.Printf("Scanned: %d\n", result.Scanned)
			fmt.Printf("Extracted: %d\n", result.Extracted)
			fmt.Printf("Written: %d\n", result.Written)
			fmt.Printf("Batches: %d\n", result.Batches)
			fmt.Printf("Time: %.3fs\n", result.Time.Seconds())
		}

		s := syncer.Create().
			SetConfig(app.Config()).
			SetDatabase(app.Database()).
			SetDrive(app.Drive()).
			SetConcurrency(concurrency).
			SetBatchSize(batchSize)

		s.OnComplete(onComplete)

		s.OnError(func(path string, err error) {
			fmt.Printf("%s\n", red.Render(path))
			fmt.Printf("Error: %v\n", err)
		})

		if !detail {
			fmt.Printf("Starting sync with concurrency %d and batch size %d\n", concurrency, batchSize)
		}

		if detail {
			s.OnSuccess(func(path string) {
				fmt.Printf("%s\n", green.Render(path))
			})
		}

		s.All()
	},
}

func init() {
	rootCmd.AddCommand(syncAllCmd)
	syncAllCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent workers")
	syncAllCmd.Flags().IntP("batch-size", "b", 100, "Number of entries per batch")
	syncAllCmd.Flags().BoolP("detail", "d", false, "Show detailed output")
}
